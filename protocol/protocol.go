package protocol

import (
	"errors"
	"log"
	"net"
	"os/exec"

	//"github.com/davecgh/go-spew/spew"

	"github.com/gorilla/websocket"
	"github.com/invctrl/hijack/client/config"
	"github.com/invctrl/hijack/client/dial"
	"github.com/navigaid/gotty/localcmd"
	"github.com/navigaid/gotty/utils"
	"github.com/navigaid/gotty/wetty"
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
	Id string
}

type ConnResponse struct {
	Message string
}

func (c *Conn) New(req ConnRequest, res *ConnResponse) error {
	log.Println("Conn.New called with", req)

	if req.Id == "" {
		return errors.New("id cannot be empty")
	}
	// log.Println(config.Default)
	res.Message = "OK"
	conn := dial.Dial(config.Default)
	dial.HandshakeAppend(conn)
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
	Id string
}

type WsConnResponse struct {
	Message string
}

func (c *WsConn) New(req WsConnRequest, res *WsConnResponse) error {
	log.Println("WsConn.New called with", req)

	if req.Id == "" {
		return errors.New("id cannot be empty")
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
