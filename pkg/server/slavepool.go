package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"
	"time"

	//"github.com/davecgh/go-spew/spew"
	"github.com/fatih/pool"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/btwiuse/gods/maps/linkedhashmap"
	"github.com/btwiuse/pretty"
	"google.golang.org/grpc"

	"github.com/btwiuse/invctrl/header"
	"github.com/btwiuse/invctrl/pkg/api"
	rpcimpl "github.com/btwiuse/invctrl/pkg/api/rpc/impl"
	"github.com/btwiuse/invctrl/wrap"
)

type Slave struct {
	UUID       uuid.UUID
	LocalAddr  net.Addr
	RemoteAddr net.Addr
	RPC        *rpc.Client
	BidiStreamClient       api.BidiStreamClient
	Info       string
	Conns      map[string]net.Conn
	WsConns    map[string]*websocket.Conn
	Pool       pool.Pool
}

type SlavePool struct {
	Slaves *linkedhashmap.Map
	Current *Slave
	Latest  *Slave
}

func NewSlavePool() *SlavePool {
	return &SlavePool{
		Slaves: linkedhashmap.New(),
	}
}

var GlobalSlavePool = NewSlavePool()

func (p *SlavePool) Del(uuid string) {
	p.Slaves.Remove(uuid)
	if (p.Current != nil) && (p.Current.UUID.String() == uuid) {
		p.Current = nil //new(Slave)
	}
}

func (p *SlavePool) GetRandom() *Slave {
	rand.Seed(time.Now().UnixNano())
	iter := p.Slaves.Iterator()
	iter.Begin()
	/*
		n := p.Slaves.Size()
		r := rand.Intn(n)
		log.Println(r, n)
		for i := 0; i <= r; i++ {
	*/
	for i := 0; i <= rand.Intn(p.Slaves.Size()); i++ {
		iter.Next()
	}
	return iter.Value().(*Slave)
}

func (p *SlavePool) Get(uuid string) *Slave {
	v, _ := p.Slaves.Get(uuid)
	return v.(*Slave)
}

func (p *SlavePool) Add(slave *Slave) {
	p.Slaves.Put(slave.UUID.String(), slave)
	p.Latest = slave
}

func (p *SlavePool) Dump() {
	log.Println("[slave pool]")
	isCurrent := func(uuid string) string {
		if (p.Current != nil) && (p.Current.UUID.String() == uuid) {
			return "*"
		}
		return " "
	}
	for i, v := range p.Slaves.Values() {
		slave := v.(*Slave)
		uuid := p.Slaves.Keys()[i].(string)
		fmt.Println(
			fmt.Sprintf("[%s]", strconv.Itoa(i+1)),
			isCurrent(uuid),
			uuid,
			"ssh ubuntu@"+strings.Split(slave.RemoteAddr.String(), ":")[0],
			slave.Info,
		)
	}
}

func (p *SlavePool) Has(uuid string) bool {
	_, found := p.Slaves.Get(uuid)
	return found
}

func NewSlave(w http.ResponseWriter) (*Slave, error) {
	slave := new(Slave) // using named return causes panic, why?
	slave.UUID = uuid.New()
	slave.WsConns = make(map[string]*websocket.Conn)
	slave.Conns = make(map[string]net.Conn)

	conn, err := wrap.WrapConn(w.(http.Hijacker).Hijack())
	if err != nil {
		return nil, err
	}

	slave.RemoteAddr = conn.RemoteAddr()
	slave.LocalAddr = conn.LocalAddr()

	sheader := &header.MasterHeader{}
	cheader := &header.SlaveHeader{}
	// here we don't cate about decoder.Buffered
	// we can pretty much assume it is empty cuz after the slave send it's header
	// to server, it will wait server's confirmation, during this period nothing will
	// be sent from the slave. so once the server receives the complete header, nothing
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
		slave := GlobalSlavePool.Get(sheader.Id)
		slave.Conns[sheader.Nonce] = conn
		log.Println(slave.UUID, "conns:", slave.Conns)
		return nil, fmt.Errorf("append to %s", sheader.Id)
	}

	cheader.Id = slave.UUID.String()
	conn.Write([]byte(pretty.JsonString(cheader)))
	slave.Info = pretty.JSONString(sheader)

	slave.RPC = NewRPCClient(conn, onclose(slave))

	log.Println("connected:", slave.UUID, slave.RemoteAddr)

	GlobalSlavePool.Add(slave)
	factory := func() (net.Conn, error) {
		log.Println("infactory")
		nonce := uuid.New().String()
		id := slave.UUID.String()
		req := rpcimpl.ConnRequest{
			Id:    id,
			Nonce: nonce,
		}
		resp := new(rpcimpl.ConnResponse)
		done := make(chan *rpc.Call, 1)
		slave.RPC.Go("Conn.New", req, resp, done)
		<-done
		conn, ok := slave.Conns[nonce]
		if !ok {
			return nil, fmt.Errorf("slave nonce doesn't exist: %s", nonce)
		}
		delete(slave.Conns, nonce)
		return conn, nil
	}
	log.Println("pool.NewChannelPool")
	slave.Pool, err = pool.NewChannelPool(5, 30, factory)
	if err != nil {
		return nil, err
	}

	log.Println("new grpc!!!")
	{
		nonce := uuid.New().String()
		id := slave.UUID.String()
		req := rpcimpl.GRPCConnRequest{
			Id:    id,
			Nonce: nonce,
		}
		resp := new(rpcimpl.GRPCConnResponse)
		done := make(chan *rpc.Call, 1)
		slave.RPC.Go("GRPCConn.New", req, resp, done)
		<-done
		time.Sleep(time.Second / 3)
		conn, ok := slave.Conns[nonce]
		if !ok {
			return nil, fmt.Errorf("slave nonce doesn't exist: %s", nonce)
		}
		delete(slave.Conns, nonce)
		cc, err := toGrpcClientConn(conn)
		if err != nil {
			return nil, err
		}
		slave.BidiStreamClient = api.NewBidiStreamClient(cc)
	}

	GlobalSlavePool.Dump()
	return slave, nil
}

// onclose is called when slave goes offline
// slave.UUID, slave.RemoteAddr, slave.Info
func onclose(slave *Slave) func() {
	return func() {
		GlobalSlavePool.Dump()
		log.Println("disconnected:", slave.UUID, slave.RemoteAddr, slave.Info)
		GlobalSlavePool.Del(slave.UUID.String())
	}
}

// we use NewRPCClient over rpc.NewClient(conn)
// so we can remove slave from pool immediately when it is disconnected

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

func toGrpcClientConn(c net.Conn) (*grpc.ClientConn, error) {
	return grpc.Dial(
		"",
		[]grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithContextDialer(
				func(ctx context.Context, s string) (net.Conn, error) {
					return c, nil
				}),
		}...,
	)
}
