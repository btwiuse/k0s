package utils

import (
	"io/ioutil"
	"log"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

/* http://www.gorillatoolkit.org/pkg/websocket
Connections support one concurrent reader and one concurrent writer.

Applications are responsible for ensuring that no more than one goroutine calls the write methods (NextWriter, SetWriteDeadline, WriteMessage, WriteJSON, EnableWriteCompression, SetCompressionLevel) concurrently and that no more than one goroutine calls the read methods (NextReader, SetReadDeadline, ReadMessage, ReadJSON, SetPongHandler, SetPingHandler) concurrently.
*/

func NetConn(c *websocket.Conn) net.Conn {
	return &netConn{
		Conn: c,
	}
}

// netConn makes a io.ReadWriteCloser from websocket.Conn, implementing the wetty.Master interface
// it is fed to wetty.New to create a WeTTY, bridging the websocket.Conn and local command
type netConn struct {
	*websocket.Conn
	mu sync.Mutex
}

func (wsw *netConn) Write(p []byte) (n int, err error) {
	wsw.mu.Lock()
	defer wsw.mu.Unlock()
	writer, err := wsw.Conn.NextWriter(websocket.BinaryMessage)
	if err != nil {
		return 0, err
	}
	defer writer.Close()
	return writer.Write(p)
}

func (wsw *netConn) Read(buf []byte) (int, error) {
	for {
		msgType, reader, err := wsw.Conn.NextReader()
		if err != nil {
			return 0, err
		}
		if msgType != websocket.BinaryMessage {
			log.Println("not BinaryMessage:")
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

func (wsw *netConn) Close() error {
	return wsw.Conn.Close()
}

func (c *netConn) SetDeadline(t time.Time) (err error) {
	err = c.SetWriteDeadline(t)
	if err != nil {
		return
	}
	err = c.SetReadDeadline(t)
	return
}
