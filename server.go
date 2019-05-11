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

type Pool map[string]*Client

var ClientPool = make(Pool)

func (p Pool) Del(uuid string) {
	delete(p, uuid)
}

func (p Pool) Add(client *Client) {
	p[client.UUID] = client
}

func (p Pool) Dump() {
	log.Println("[client pool]")
	for uuid, client := range p {
		fmt.Println(uuid, client.Conn.RemoteAddr())
	}
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
		defer log.Println("disconnected:", client.Conn.RemoteAddr())
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

	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-quit:
			return
		default:
			if scanner.Scan() {
				line := scanner.Text() + "\n"
				conn.Write([]byte(line))
			} else {
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/", hijacker)
	log.Println("listening on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
