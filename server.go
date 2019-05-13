package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/navigaid/gods/maps/linkedhashmap"
	"github.com/navigaid/pretty"
	"gopkg.in/readline.v1"
)

type Client struct {
	UUID string
	Conn net.Conn
	Quit chan struct{}
	Info string
}

func NewClient(uuid string, conn net.Conn, quit chan struct{}) *Client {
	return &Client{
		UUID: uuid,
		Conn: conn,
		Quit: quit,
	}
}

type Pool struct {
	Clients *linkedhashmap.Map // map[string]*Client
	Current *Client
}

func NewPool() *Pool {
	return &Pool{
		Clients: linkedhashmap.New(), //make(map[string]*Client),
	}
}

var ClientPool = NewPool()

func (p *Pool) Del(uuid string) {
	// delete(p.Clients, uuid)
	p.Clients.Remove(uuid)
	if (p.Current != nil) && (p.Current.UUID == uuid) {
		p.Current = nil //new(Client)
	}
}

func (p *Pool) Get(uuid string) *Client {
	// return p.Clients[uuid]
	v, _ := p.Clients.Get(uuid)
	return v.(*Client)
}

func (p *Pool) Add(client *Client) {
	// p.Clients[client.UUID] = client
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
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func hijacker(w http.ResponseWriter, r *http.Request) {
	uuid := uuid.New().String()
	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		log.Println(nil)
	}

	quit := make(chan struct{})
	client := NewClient(uuid, conn, quit)

	var v map[string]interface{}
	decoder := json.NewDecoder(conn)
	decoder.Decode(&v)
	client.Info = pretty.JSONString(v)

	log.Println("connected:", uuid, conn.RemoteAddr(), client.Info)

	ClientPool.Add(client)
	ClientPool.Dump()
	copy := func(dst io.Writer, src io.Reader) {
		defer ClientPool.Dump()
		defer log.Println("disconnected:", uuid, client.Conn.RemoteAddr(), client.Info)
		defer close(client.Quit)
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
	go copy(os.Stdout, io.MultiReader(decoder.Buffered(), conn))
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

			ClientPool.Current.Conn.Write([]byte(line + "\n"))

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
