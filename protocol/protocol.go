package protocol

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os/exec"

	"github.com/gorilla/websocket"
	"github.com/btwiuse/gotty/localcmd"
	"github.com/btwiuse/gotty/utils"
	"github.com/btwiuse/gotty/wetty"
	"google.golang.org/grpc"

	"github.com/btwiuse/invctrl/client/config"
	"github.com/btwiuse/invctrl/client/dial"
	"github.com/btwiuse/invctrl/pkg/api"
)

type Request struct {
	Command string
}

type Response struct {
	Message string
}

type Hello struct{}

func (h *Hello) Execute(req Request, res *Response) error {
	log.Println("Hello called with", req)
	if req.Command == "" {
		return errors.New("command cannot be empty")
	}
	res.Message = "Hello " + req.Command
	return nil
}

type Bash struct{}

func (h *Bash) Execute(req Request, res *Response) error {
	log.Println("Bash called with", req)

	if req.Command == "" {
		return errors.New("command cannot be empty")
	}
	cmd := exec.Command("/bin/bash", "-c", req.Command)
	out, err := cmd.CombinedOutput()
	res.Message = string(out)
	return err
}

type Conn struct{}

type ConnRequest struct {
	Id    string
	Nonce string
}

type ConnResponse struct {
	Message string
}

func (c *Conn) New(req ConnRequest, res *ConnResponse) error {
	log.Println("Conn.New called with", req)

	if req.Id == "" {
		return errors.New("id cannot be empty")
	}
	if req.Nonce == "" {
		return errors.New("nonce cannot be empty")
	}
	// log.Println(config.Default)
	res.Message = "OK"
	conn := dial.Dial(config.Default)
	dial.HandshakeAppend(conn, req.Nonce)
	println("dial.HandshakeAppend")
	go serveHTTP(conn)
	return nil
}

func serveHTTP(conn net.Conn) {
	conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Length: 76\r\nContent-Type: text/plain; charset=utf-8\r\nDate: Wed, 19 Jul 1972 19:00:00 GMT\r\n\r\nGo is a general-purpose language designed with systems programming in mind.\n"))
	// conn.Close()
}

type WsConn struct{}

type WsConnRequest struct {
	Id    string
	Nonce string
}

type WsConnResponse struct {
	Message string
}

func (c *WsConn) New(req WsConnRequest, res *WsConnResponse) error {
	log.Println("WsConn.New called with", req)

	if req.Id == "" {
		return errors.New("id cannot be empty")
	}
	if req.Nonce == "" {
		return errors.New("nonce cannot be empty")
	}
	// log.Println(config.Default)
	res.Message = "OK"
	conn := dial.WsDial(config.Default)
	factory := &localcmd.Factory{
		Args: []string{"/bin/bash"},
		//Args: []string{"neofetch"},
		//Args: []string{"htop"},
		//Args: []string{"tmux"},
		//Args: []string{"seq", "100"},
		//Args: []string{"/home/aaron/go/bin/z", "tick", "100"},
	}
	rw := &utils.WsWrapper{conn}
	rw.Write([]byte(string(req.Id)))
	rw.Write([]byte(string(req.Nonce)))
	go serveWS(conn, factory)
	return nil
}

func serveWS(conn *websocket.Conn, factory *localcmd.Factory) {
	var master wetty.Master = &utils.WsWrapper{conn}
	var slave wetty.Slave
	var err error
	slave, err = factory.New()
	if err != nil {
		log.Println(err)
		return
	}
	if err := wetty.NewMSPair(master, slave).Pipe(); err != nil {
		log.Println(err)
	}
}

type Rootfs struct{}

type RootfsRequest struct {
	Request []byte
}

type RootfsResponse struct {
	Response []byte
}

func (*Rootfs) New(req RootfsRequest, res *RootfsResponse) error {
	log.Println("Rootfs.New called with", string(req.Request))
	w := httptest.NewRecorder()
	r, err := http.ReadRequest(bufio.NewReader(bytes.NewBuffer(req.Request)))
	if err != nil {
		return err
	}
	http.FileServer(http.Dir("/")).ServeHTTP(w, r)
	res.Response, err = httputil.DumpResponse(w.Result(), true)
	if err != nil {
		return err
	}
	return nil
}

type Echo struct{}

type EchoRequest struct {
	Payload string
}

type EchoResponse struct {
	Payload string
}

func (*Echo) New(req EchoRequest, res *EchoResponse) error {
	log.Println("Echo.New called with", req.Payload)
	res.Payload = req.Payload
	return nil
}

type GRPCConn struct{}

type GRPCConnRequest struct {
	Id    string
	Nonce string
}

type GRPCConnResponse struct {
	Message string
}

func (*GRPCConn) New(req ConnRequest, res *ConnResponse) error {
	log.Println("GRPCConn.New called with", req)

	if req.Id == "" {
		return errors.New("id cannot be empty")
	}
	if req.Nonce == "" {
		return errors.New("nonce cannot be empty")
	}
	// log.Println(config.Default)
	res.Message = "OK"
	conn := dial.Dial(config.Default)
	dial.HandshakeAppend(conn, req.Nonce)
	println("dial.HandshakeAppend")
	go serveGRPC(conn)
	return nil
}

func serveGRPC(grpcSide net.Conn) error {
	l := &singleListener{grpcSide}
	grpcServer := grpc.NewServer()
	api.RegisterApiServer(grpcServer, &apiService{})
	return grpcServer.Serve(l)
}

type singleListener struct {
	conn net.Conn
}

/*
type Listener interface {
        // Accept waits for and returns the next connection to the listener.
        Accept() (Conn, error)

        // Close closes the listener.
        // Any blocked Accept operations will be unblocked and return errors.
        Close() error

        // Addr returns the listener's network address.
        Addr() Addr
}
*/

// singleListener implements the net.Listener interface
func (s *singleListener) Accept() (net.Conn, error) {
	log.Println("Accept")
	if c := s.conn; c != nil {
		s.conn = nil
		return c, nil
	}
	return nil, io.EOF
}

func (s *singleListener) Close() error {
	return nil
}

func (s *singleListener) Addr() net.Addr {
	return s.conn.LocalAddr()
}

// apiService is an api.Service
type apiService struct {
}

func (s *apiService) Hello(ctx context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {
	log.Println("apiService.Hello called with", req.Name)
	return &api.HelloResponse{Message: "Hello " + req.GetName()}, nil
}
