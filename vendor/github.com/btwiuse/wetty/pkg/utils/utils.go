package utils

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"github.com/gorilla/websocket"
)

/* http://www.gorillatoolkit.org/pkg/websocket
Connections support one concurrent reader and one concurrent writer.

Applications are responsible for ensuring that no more than one goroutine calls the write methods (NextWriter, SetWriteDeadline, WriteMessage, WriteJSON, EnableWriteCompression, SetCompressionLevel) concurrently and that no more than one goroutine calls the read methods (NextReader, SetReadDeadline, ReadMessage, ReadJSON, SetPongHandler, SetPingHandler) concurrently.
*/

func WsConnToReadWriter(c *websocket.Conn) *WsWrapper {
	ret := new(WsWrapper)
	ret.Conn = c
	return ret
}

// WsWrapper makes a io.ReadWriter from websocket.Conn, implementing the wetty.Master interface
// it is fed to wetty.New to create a WeTTY, bridging the websocket.Conn and local command
type WsWrapper struct {
	*websocket.Conn
	mu sync.Mutex
}

func (wsw *WsWrapper) Write(p []byte) (n int, err error) {
	wsw.mu.Lock()
	defer wsw.mu.Unlock()
	writer, err := wsw.Conn.NextWriter(websocket.BinaryMessage)
	if err != nil {
		return 0, err
	}
	defer writer.Close()
	return writer.Write(p)
}

func (wsw *WsWrapper) Read(buf []byte) (int, error) {
	for {
		msgType, reader, err := wsw.Conn.NextReader()
		if err != nil {
			return 0, err
		}
		if msgType != websocket.BinaryMessage {
			continue
		}

		msg, err := ioutil.ReadAll(reader)
		if err != nil {
			return 0, err
		}

		copy(buf, msg)

		n := len(msg)
		if n > len(buf) {
			n = len(buf)
		}

		return n, nil
	}
}

func (wsw *WsWrapper) Close() error {
	return wsw.Conn.Close()
}

// ReadWriter stores pointers to a Reader and a Writer.
// It implements io.ReadWriter automatically
type ReadWriter struct {
	io.Reader
	io.Writer
}

// single listener converts/upgrades the current tcp connection into grpc
// ============================= gender changer impl
type SingleListener struct {
	net.Conn
}

// SingleListener implements the net.Listener interface
func (s *SingleListener) Accept() (net.Conn, error) {
	if s.Conn != nil {
		log.Println("Gender Change: TCP Client -> GRPC Server")
		c := s.Conn
		s.Conn = nil
		return c, nil
	}
	return nil, io.EOF
}

func (s *SingleListener) Close() error {
	return nil
}

func (s *SingleListener) Addr() net.Addr {
	return s.Conn.LocalAddr()
}
