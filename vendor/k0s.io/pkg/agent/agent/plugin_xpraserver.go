package agent

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"os"

	"k0s.io"
	types "k0s.io/pkg/agent"
	"k0s.io/pkg/api"
	"k0s.io/pkg/log"
	"nhooyr.io/websocket"
)

func init() { Tunnels[api.Xpra] = StartXpraServer }

func StartXpraServer(c types.Config) chan net.Conn {
	xpraListener := NewLys()
	go xpraServe(xpraListener)
	return xpraListener.Conns
}

func wsdial(u string) (*websocket.Conn, net.Conn, error) {
	var (
		h = http.Header{}
		t = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
		opts = &websocket.DialOptions{
			HTTPClient:      t,
			HTTPHeader:      h,
			Subprotocols:    []string{"binary"},
			CompressionMode: websocket.CompressionContextTakeover,
		}
	)

	wsconn, _, err := websocket.Dial(context.Background(), u, opts)
	if err != nil {
		return nil, nil, err
	}
	wsconn.SetReadLimit(k0s.MAX_WS_MESSAGE)
	return wsconn, websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary), nil
}

func xpraHandle(c net.Conn) {
	u := envXpra()

	log.Println("XPRA: ", u, c.RemoteAddr())

	_, dnetconn, err := wsdial(u)
	if err != nil {
		log.Println(err)
		c.Close()
		return
	}

	// go io.Copy(c, d)
	// io.Copy(d, c)
	// _ = d
	// io.Copy(os.Stderr, c)
	// go io.Copy(os.Stderr, d)
	// b2 := make([]byte, 640 * 1024)
	/*
		b := make([]byte, 160 * 1024)
		go io.CopyBuffer(c, dnetconn, b) // local to remote
	*/

	go func() {
		if _, err := io.CopyBuffer(c, dnetconn, make([]byte, k0s.MAX_WS_MESSAGE)); err != nil {
			log.Println(err)
		}
	}()

	/*
		go func(){
			for {
				_, buf, err := dwsconn.Read(context.Background())
				if err != nil {
					log.Println(err)
					break
				}
				n, err := c.Write(buf)
				if err != nil {
					log.Println(err)
					break
				}
				// log.Println("wrote", len(buf), n)
			}
		}() // local to remote
	*/

	io.Copy(dnetconn, c) // remote to local
}

func xpraServe(l net.Listener) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		go func() {
			defer c.Close()
			xpraHandle(c)
		}()
	}
}

func envXpra() string {
	ep, ok := os.LookupEnv("XPRA")
	if !ok {
		return ""
	}
	return ep
}
