package hub

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/btwiuse/wetty/pkg/msg"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"k0s.io/conntroll/pkg/api"
	types "k0s.io/conntroll/pkg/hub"
	"k0s.io/conntroll/pkg/wrap"
	"nhooyr.io/websocket"
)

// new wire format
func terminalRelay(ag types.Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wsc, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			InsecureSkipVerify: true,
			Subprotocols:       []string{"wetty"},
		})
		if err != nil {
			log.Println(err)
			return
		}
		wsconn := websocket.NetConn(context.Background(), wsc, websocket.MessageBinary)
		defer wsconn.Close()

		conn := ag.NewTerminal()
		defer conn.Close()

		go io.Copy(wsconn, conn)
		io.Copy(conn, wsconn)
	}
}

// deprecated in favor of terminalRelay
func wsRelay(ag types.Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wsconn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			InsecureSkipVerify: true,
			Subprotocols:       []string{"wetty"},
		})
		if err != nil {
			log.Println(err)
			return
		}
		conn := websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary)

		session := ag.NewSession()
		defer session.Close()
		sessionSendClient, err := session.Send(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		// common error: ws transport is closing
		log.Println(pipe(conn, sessionSendClient))
	}
}

// (through chan Message{Type, Body} instead of interface)
func pipe(ws io.ReadWriteCloser, session api.Session_SendClient) error {
	defer ws.Close()
	g, ctx := errgroup.WithContext(context.TODO())
	_ = ctx
	g.Go(func() error {
		log.Println("pipe: client(ws) => session(grpc)")
		// TODO: io.Copy(session, ws), CopyBuffer, session.ReadFrom
		buf := make([]byte, 1<<12) // maximum input message is 4096 bytes
		for {
			n, err := ws.Read(buf)
			if err != nil {
				return err
			}
			msg := &api.Message{Type: msg.Type(buf[0]), Body: buf[1:n]}
			err = session.Send(msg)
			if err != nil {
				return err
			}
		}
		return nil
	})

	g.Go(func() error {
		log.Println("pipe: client(ws) <= session(grpc)")
		// TODO: io.Copy(ws, session), CopyBuffer, session.WriteTo
		for {
			resp, err := session.Recv()
			if err != nil {
				return err
			}
			_, err = ws.Write(append([]byte{byte(resp.Type)}, resp.Body...))
			if err != nil {
				return err
				break
			}
		}
		return nil
	})

	return g.Wait()
}

func fsRelay(ag types.Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			vars = mux.Vars(r)
			id   = vars["id"]
			path = strings.TrimPrefix(r.RequestURI, "/api/agent/"+id+"/rootfs")
		)
		r.RequestURI = path

		reqbuf, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
			return
		}

		conn, err := wrap.Hijack(w)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		fsConn := ag.NewFS()
		defer fsConn.Close()

		go func() {
			io.Copy(fsConn, bytes.NewBuffer(reqbuf))
		}()
		io.Copy(conn, fsConn)
	}
}

func metricsRelay(ag types.Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			vars = mux.Vars(r)
			id   = vars["id"]
			path = strings.TrimPrefix(r.RequestURI, "/api/agent/"+id)
		)
		r.RequestURI = path

		reqbuf, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
			return
		}

		conn, err := wrap.Hijack(w)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		metricsConn := ag.NewMetrics()
		defer metricsConn.Close()

		go func() {
			io.Copy(metricsConn, bytes.NewBuffer(reqbuf))
		}()
		io.Copy(conn, metricsConn)
	}
}

func socks5Relay(ag types.Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wsconn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			InsecureSkipVerify: true,
		})
		if err != nil {
			log.Println(err)
			return
		}
		conn := websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary)

		socks5Conn := ag.NewSocks5()
		defer socks5Conn.Close()

		go func() {
			_, err := io.Copy(conn, socks5Conn)
			if err != nil {
				log.Println(err)
				return
			}
		}()

		_, err = io.Copy(socks5Conn, conn)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
