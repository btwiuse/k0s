package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/invctrl/hijack/protocol"
	"github.com/invctrl/hijack/wrap"
	"github.com/navigaid/gods/maps/linkedhashmap"
	"github.com/navigaid/pretty"
	"gopkg.in/readline.v1"
)

type Client struct {
	UUID string
	Conn net.Conn
	RPC  *rpc.Client
	Info string
}

type Pool struct {
	Clients *linkedhashmap.Map
	Current *Client
}

func NewPool() *Pool {
	return &Pool{
		Clients: linkedhashmap.New(),
	}
}

var ClientPool = NewPool()

func (p *Pool) Del(uuid string) {
	p.Clients.Remove(uuid)
	if (p.Current != nil) && (p.Current.UUID == uuid) {
		p.Current = nil //new(Client)
	}
}

func (p *Pool) Get(uuid string) *Client {
	v, _ := p.Clients.Get(uuid)
	return v.(*Client)
}

func (p *Pool) Add(client *Client) {
	p.Clients.Put(client.UUID, client)
}

func (p *Pool) Dump() {
	log.Println("[client pool]")
	isCurrent := func(uuid string) string {
		if (p.Current != nil) && (p.Current.UUID == uuid) {
			return "*"
		}
		return " "
	}
	for i, v := range p.Clients.Values() {
		client := v.(*Client)
		uuid := p.Clients.Keys()[i].(string)
		fmt.Println(
			fmt.Sprintf("[%s]", strconv.Itoa(i+1)),
			isCurrent(uuid),
			uuid,
			"ssh ubuntu@"+strings.Split(client.Conn.RemoteAddr().String(), ":")[0],
			client.Info,
		)
	}
}

func (p *Pool) Has(uuid string) bool {
	_, found := p.Clients.Get(uuid)
	return found
}

func lojacker(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(http.StatusText(http.StatusOK)))
	isCurrent := func(uuid string) string {
		if (ClientPool.Current != nil) && (ClientPool.Current.UUID == uuid) {
			return "*"
		}
		return " "
	}
	w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	for i, v := range ClientPool.Clients.Values() {
		client := v.(*Client)
		uuid := ClientPool.Clients.Keys()[i].(string)
		fmt.Fprintln(w,
			fmt.Sprintf("[%s]", strconv.Itoa(i+1)),
			isCurrent(uuid),
			uuid,
			"ssh ubuntu@"+strings.Split(client.Conn.RemoteAddr().String(), ":")[0],
			client.Info,
		)
	}
}

func getHeader(r io.Reader) (string, io.Reader, error) {
	var v interface{}
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&v); err != nil {
		return "", nil, err
	}
	return pretty.JSONString(v), decoder.Buffered(), nil
}

func newClient(w http.ResponseWriter) (client *Client, err error) {
	uuid := uuid.New().String()

	conn, hibuf, err := w.(http.Hijacker).Hijack()
	if err != nil {
		return nil, err
	}

	header, rest, err := getHeader(io.MultiReader(hibuf, conn))
	if err != nil {
		conn.Write([]byte("NO"))
		return nil, err
	}
	conn.Write([]byte("OK"))

	rpcClient := NewRPC(
		io.MultiReader(hibuf, rest),
		conn,
		callback(uuid, conn.RemoteAddr().String()),
	)

	return &Client{
		UUID: uuid,
		Conn: conn,
		RPC:  rpcClient,
		Info: header,
	}, nil
}

func callback(uuid string, raddr string) func() {
	return func() {
		ClientPool.Dump()
		log.Println("disconnected:", uuid, raddr)
		ClientPool.Del(uuid)
	}
}

// we use NewRPC over rpc.NewClient(conn)
// so we can remove client from pool immediately when it is disconnected
func NewRPC(mr io.Reader, conn io.ReadWriteCloser, callback func()) *rpc.Client {
	copy := func(dst io.Writer, src io.Reader) {
		if _, err := io.Copy(dst, src); err != nil {
			log.Println(err)
		}
		callback()
	}
	pr, pw := io.Pipe()
	go copy(pw, io.MultiReader(mr, conn))
	return rpc.NewClient(wrap.WrapReadWriteCloser(pr, conn))
}

func hijacker(w http.ResponseWriter, r *http.Request) {
	client, err := newClient(w)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("connected:", client.UUID, client.Conn.RemoteAddr())
	ClientPool.Add(client)
	ClientPool.Dump()
}

func input() {
	for {
		log.Println("ready to accept input!")
		rl, err := readline.NewEx(&readline.Config{
			HistoryFile:         "/tmp/readline.tmp",
			ForceUseInteractive: true,
			// InterruptPrompt:     "exit?",
		})
		if err != nil {
			panic(err)
		}
		defer rl.Close()
		fmt.Println("Welcome to InvCtrl!!!")
		promptNum := 1
	INNER:
		for {
			rl.SetPrompt(fmt.Sprintf("In[%d]:= ", promptNum))
			line, err := rl.Readline()
			switch err {
			case nil: // NOP
			case io.EOF:
				fmt.Println("bye")
				break INNER
			case readline.ErrInterrupt:
				fmt.Println("try Exit, or Quit")
			default:
				fmt.Println(err)
			}

			bash := func(line string, client *Client) {
				req := protocol.Request{
					Command: line,
				}
				resp := new(protocol.Response)
				err := client.RPC.Call("Bash.Execute", req, resp)
				if err != nil {
					log.Println(resp.Message, err)
					return
				}
				log.Println("rpc message received:\n\n", resp.Message)
			}
			switch {
			case strings.HasPrefix(line, "!map "):
				line = strings.TrimPrefix(line, "!map ")
				for _, v := range ClientPool.Clients.Values() {
					client := v.(*Client)
					//client.Conn.Write([]byte(line + "\n"))
					go bash(line, client)
				}
				continue
			case line == "":
				continue
			case line == "Exit", line == "Quit":
				os.Exit(0)
			case line == "Ls":
				ClientPool.Dump()
				continue
			case ClientPool.Has(line):
				ClientPool.Current = ClientPool.Get(line)
				log.Println("current client:", ClientPool.Current.UUID)
				continue
			default:
				if ClientPool.Current == nil {
					fmt.Println("[INFO] Your current client is empty. Enter the uuid to the client you want to talk to first:")
					continue
				}
			}

			go bash(line, ClientPool.Current)

			promptNum += 1
		}

		log.Println("stdin input closed")
	}
}

func hilo(hijacker, lojacker func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pretty.JSON(r.Header)
		if r.Header.Get(http.CanonicalHeaderKey("Hijack")) == "true" {
			log.Println("hijack")
			hijacker(w, r)
			return
		}
		log.Println("lojack")
		lojacker(w, r)
	}
}

func main() {
	http.HandleFunc("/", hilo(hijacker, lojacker))
	log.Println("listening on http://localhost:8000")
	go input()
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
