package hub

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/gorilla/mux"
	types "k0s.io/k0s/pkg/hub"
	"k0s.io/k0s/pkg/wrap"
	"nhooyr.io/websocket"
)

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

		go io.Copy(conn, wsconn)
		io.Copy(wsconn, conn)
	}
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

func redirRelay(ag types.Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wsconn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			InsecureSkipVerify: true,
		})
		if err != nil {
			log.Println(err)
			return
		}
		conn := websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary)

		redirConn := ag.NewRedir()
		defer redirConn.Close()

		go func() {
			_, err := io.Copy(conn, redirConn)
			if err != nil {
				log.Println(err)
				return
			}
		}()

		_, err = io.Copy(redirConn, conn)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
