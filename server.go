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
	//"time"

	"github.com/google/uuid"
	"github.com/invctrl/hijack/protocol"
	"github.com/navigaid/gods/maps/linkedhashmap"
	"github.com/navigaid/pretty"
	"gopkg.in/readline.v1"
)

type Client struct {
	UUID string
	Conn net.Conn
	RPC  *rpc.Client
	Quit chan struct{}
	Info string
}

func NewClient(uuid string, conn net.Conn, quit chan struct{}, rpc *rpc.Client) *Client {
	return &Client{
		UUID: uuid,
		Conn: conn,
		Quit: quit,
		RPC:  rpc,
	}
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

func hijacker(w http.ResponseWriter, r *http.Request) {
	uuid := uuid.New().String()
	conn, hibuf, err := w.(http.Hijacker).Hijack()
	if err != nil {
		log.Println(nil)
	}
	quit := make(chan struct{})

	v := make(map[string]string)
	decoder := json.NewDecoder(io.MultiReader(hibuf, conn))
	if err := decoder.Decode(&v); err != nil {
		log.Println(err)
		conn.Write([]byte("NO"))
		conn.Close()
		return
	}
	conn.Write([]byte("OK"))
	header := pretty.JSONString(v)

	log.Println("connected:", uuid, conn.RemoteAddr(), header)

	pr, pw := io.Pipe()
	copy := func(dst io.Writer, src io.Reader) {
		defer ClientPool.Dump()
		defer log.Println("disconnected:", uuid, conn.RemoteAddr())
		defer close(quit)
		defer ClientPool.Del(uuid)
		buf := make([]byte, 1)
		for {
			_, err := src.Read(buf)
			if err != nil {
				return
			}
			dst.Write(buf)
		}
	}
	go copy(pw, io.MultiReader(hibuf, decoder.Buffered(), conn))

	rpc.Register(new(protocol.Hello))
	rpcClient := rpc.NewClient(NewRWC(pr, conn))

	client := NewClient(uuid, conn, quit, rpcClient)
	client.Info = header

	ClientPool.Add(client)
	ClientPool.Dump()
}

type brwc struct {
	pr  io.Reader
	rwc io.ReadWriteCloser
}

func NewRWC(pr io.Reader, rwc io.ReadWriteCloser) *brwc {
	return &brwc{
		pr:  pr,
		rwc: rwc,
	}
}

func (b *brwc) Close() error {
	return b.rwc.Close()
}

func (b *brwc) Write(p []byte) (int, error) {
	return b.rwc.Write(p)
}

func (b *brwc) Read(p []byte) (int, error) {
	return io.MultiReader(b.pr, b.rwc).Read(p)
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

			switch {
			case strings.HasPrefix(line, "!map "):
				line = strings.TrimPrefix(line, "!map ")
				for _, v := range ClientPool.Clients.Values() {
					client := v.(*Client)
					client.Conn.Write([]byte(line + "\n"))
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

			if strings.HasPrefix(line, "rpc ") {
				line = strings.TrimPrefix(line, "rpc ")
				req := &protocol.Request{
					Command: line,
				}
				resp := new(protocol.Response)
				err := ClientPool.Current.RPC.Call("Hello.Execute", req, resp)
				if err != nil {
					log.Println(err)
					continue
				}
				log.Println("rpc message received:", resp.Message)
			} else {
				ClientPool.Current.Conn.Write([]byte(line + "\n"))
			}

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
