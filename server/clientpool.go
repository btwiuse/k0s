package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"

	//"github.com/davecgh/go-spew/spew"
	"github.com/fatih/pool"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/invctrl/hijack/header"
	"github.com/invctrl/hijack/protocol"
	"github.com/invctrl/hijack/wrap"
	"github.com/navigaid/gods/maps/linkedhashmap"
	"github.com/navigaid/pretty"
)

type Client struct {
	UUID       uuid.UUID
	LocalAddr  net.Addr
	RemoteAddr net.Addr
	RPC        *rpc.Client
	Info       string
	Conns      map[string]net.Conn
	WsConns    map[string]*websocket.Conn
	Pool       pool.Pool
}

type Pool struct {
	Clients *linkedhashmap.Map
	Current *Client
	Latest  *Client
}

func NewPool() *Pool {
	return &Pool{
		Clients: linkedhashmap.New(),
	}
}

var ClientPool = NewPool()

func (p *Pool) Del(uuid string) {
	p.Clients.Remove(uuid)
	if (p.Current != nil) && (p.Current.UUID.String() == uuid) {
		p.Current = nil //new(Client)
	}
}

func (p *Pool) Get(uuid string) *Client {
	v, _ := p.Clients.Get(uuid)
	return v.(*Client)
}

func (p *Pool) Add(client *Client) {
	p.Clients.Put(client.UUID.String(), client)
	p.Latest = client
}

func (p *Pool) Dump() {
	log.Println("[client pool]")
	isCurrent := func(uuid string) string {
		if (p.Current != nil) && (p.Current.UUID.String() == uuid) {
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
			"ssh ubuntu@"+strings.Split(client.RemoteAddr.String(), ":")[0],
			client.Info,
		)
	}
}

func (p *Pool) Has(uuid string) bool {
	_, found := p.Clients.Get(uuid)
	return found
}

func NewClient(w http.ResponseWriter) (*Client, error) {
	client := new(Client) // using named return causes panic, why?
	client.UUID = uuid.New()
	client.WsConns = make(map[string]*websocket.Conn)
	client.Conns = make(map[string]net.Conn)

	conn, err := wrap.WrapConn(w.(http.Hijacker).Hijack())
	if err != nil {
		return nil, err
	}

	client.RemoteAddr = conn.RemoteAddr()
	client.LocalAddr = conn.LocalAddr()

	sheader := &header.Header{}
	cheader := &header.ClientHeader{}
	// here we don't cate about decoder.Buffered
	// we can pretty much assume it is empty cuz after the client send it's header
	// to server, it will wait server's confirmation, during this period nothing will
	// be sent from the client. so once the server receives the complete header, nothing
	// is left in the buffer
	// similarly hibuf is useless after this...
	if err := json.NewDecoder(conn).Decode(&sheader); err != nil {
		cheader.Status = "NO"
		cheader.Err = err.Error()
		conn.Write([]byte(pretty.JsonString(cheader)))
		return nil, err
	}

	cheader.Status = "OK"

	if sheader.Append {
		conn.Write([]byte(pretty.JsonString(cheader)))
		client := ClientPool.Get(sheader.Id)
		client.Conns[sheader.Nonce] = conn
		log.Println(client.UUID, "conns:", client.Conns)
		return nil, fmt.Errorf("append to %s", sheader.Id)
	}

	cheader.Id = client.UUID.String()
	conn.Write([]byte(pretty.JsonString(cheader)))
	client.Info = pretty.JSONString(sheader)

	client.RPC = NewRPCClient(conn, onclose(client))

	log.Println("connected:", client.UUID, client.RemoteAddr)

	ClientPool.Add(client)
	ClientPool.Get(client.UUID.String())
	factory := func() (net.Conn, error) {
		nonce := uuid.New().String()
		id := client.UUID.String()
		req := protocol.ConnRequest{
			Id:    id,
			Nonce: nonce,
		}
		resp := new(protocol.ConnResponse)
		done := make(chan *rpc.Call, 1)
		client.RPC.Go("Conn.New", req, resp, done)
		<-done
		conn, ok := client.Conns[nonce]
		if !ok {
			return nil, fmt.Errorf("client nonce doesn't exist: %s", nonce)
		}
		delete(client.Conns, nonce)
		return conn, nil
	}
	client.Pool, err = pool.NewChannelPool(5, 30, factory)
	if err != nil {
		return nil, err
	}

	ClientPool.Dump()
	return client, nil
}

// onclose is called when client goes offline
// client.UUID, client.RemoteAddr, client.Info
func onclose(client *Client) func() {
	return func() {
		ClientPool.Dump()
		log.Println("disconnected:", client.UUID, client.RemoteAddr, client.Info)
		ClientPool.Del(client.UUID.String())
	}
}

// we use NewRPCClient over rpc.NewClient(conn)
// so we can remove client from pool immediately when it is disconnected

/*
          / io.Reader >--->copy()---> io.PipeWriter ===> io.PipeReader = intercepted io.Reader \
net.Conn  - io.Writer \                                                                        wrap.ReadWriteCloser() - rpc.NewClient - *rpc.Client
          \ io.Closer - io.WriteCloser ---------------------------------------------------------
*/
func NewRPCClient(conn io.ReadWriteCloser, onclose func()) *rpc.Client {
	copy := func(dst io.Writer, src io.Reader) {
		if _, err := io.Copy(dst, src); err != nil {
			log.Println(err)
		}
		onclose()
	}
	pr, pw := io.Pipe()
	go copy(pw, conn)
	return rpc.NewClient(wrap.WrapReadWriteCloser(pr, conn))
}
