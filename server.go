package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/google/uuid"
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
		Current: new(Client),
	}
}

var ClientPool = NewPool()

func (p *Pool) Del(uuid string) {
	delete(p.Clients, uuid)
	if (p.Current != nil) && (p.Current.UUID == uuid) {
		p.Current = new(Client)
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
		scanner := bufio.NewScanner(os.Stdin)
		log.Println("ready to accept input!")
		for scanner.Scan() {
			line := scanner.Text()
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
		}
		log.Println("stdin input closed")
	}
}

func main() {
	http.HandleFunc("/", hijacker)
	log.Println("listening on http://localhost:8000")
	go input()
	http.ListenAndServe(":8000", nil)
}
