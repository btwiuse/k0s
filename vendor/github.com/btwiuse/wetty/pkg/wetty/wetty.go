// Package wetty provides a protocol and an implementation to
// control terminals thorough networks.
package wetty

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/btwiuse/wetty/pkg/msg"
	"github.com/gorilla/websocket"
	"github.com/kr/pty"
)

// Protocols defines the name of this protocol,
// which is supposed to be used to the subprotocol of Websockt streams.
// https://stackoverflow.com/questions/37984320/why-doesnt-golang-allow-const-maps
var (
	Protocols = []string{
		"wetty",
	}
	Dialer = &websocket.Dialer{
		Subprotocols:      Protocols,
		EnableCompression: true,
	}
	Upgrader = &websocket.Upgrader{
		Subprotocols:      Protocols,
		EnableCompression: true,
		CheckOrigin:       func(r *http.Request) bool { return true },
	}
)

type (
	Client interface {
		io.Reader
		io.Writer
		io.Closer
	}

	Session interface {
		io.Reader
		io.Writer
		io.Closer
		ResizeTerminal(*pty.Winsize) error
	}

	// WeTTY bridges a PTY slave and its PTY master.
	// To support text-based streams and side channel commands such as
	// terminal resizing, WeTTY uses an original protocol.
	ClientSessionPair struct {
		client     Client  // PTY Master, which probably a connection to browser
		session    Session // PTY Slave
		bufferSize int
	}
)

///wetty.go

// New creates a new instance of WeTTY.
// masterConn is a connection to the PTY master,
// typically it's a websocket connection to a client.
// slave is a PTY slave such as a local command with a PTY.
func NewClientSessionPair(client Client, session Session) *ClientSessionPair {
	return &ClientSessionPair{
		client:     client,
		session:    session,
		bufferSize: 4096, // this means max websocket message size will be 4096 + 1(msgType)
	}
}

// slave >>>{ Output }>>> master >>> client
// partition raw output in to frames with Output header
func sessionToClient(client Client, session Session, buf []byte) error {
	for {
		n, err := session.Read(buf)
		if err != nil {
			return err
		}

		_, err = client.Write(append([]byte{byte(msg.Type_SESSION_OUTPUT)}, buf[:n]...))
		if err != nil {
			return err
		}
	}
}

// slave <<<{ Close }<<< master <<<{ ResizeTerminal, Input }<<< client
func clientToSession(session Session, client Client, buf []byte) error {
	for {
		n, err := client.Read(buf)
		if err != nil {
			return err
		}
		switch msgType := msg.Type(buf[0]); msgType {
		case msg.Type_CLIENT_INPUT: // written by client
			_, err = session.Write(buf[1:n])
			if err != nil {
				return err
			}
		case msg.Type_SESSION_RESIZE: // written by client
			sz := &pty.Winsize{}
			err = json.Unmarshal(buf[1:n], sz)
			if err != nil {
				return err
			}
			err = session.ResizeTerminal(sz)
			if err != nil {
				return err
			}
			log.Println("new sz:", sz)
		case msg.Type_SESSION_CLOSE: // written by client
			err = session.Close()
			if err != nil {
				return err
			}
			log.Println("Close, kill session")
		}
	}
}

// when to call CSPair, who calls it?
// from the perspective of master/server
func (ms *ClientSessionPair) Pipe() error {
	errs := make(chan error, 2)

	go func() {
		buf := make([]byte, ms.bufferSize)
		errs <- sessionToClient(ms.client, ms.session, buf)
	}()

	go func() {
		buf := make([]byte, ms.bufferSize)
		errs <- clientToSession(ms.session, ms.client, buf)
	}()

	return <-errs
}
