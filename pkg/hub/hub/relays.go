package hub

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"

	"github.com/btwiuse/wetty/pkg/msg"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"k0s.io/conntroll/pkg/api"
	types "k0s.io/conntroll/pkg/hub"
	"k0s.io/conntroll/pkg/wrap"
	"nhooyr.io/websocket"
)

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
		defer conn.Close()

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

		conn, err := wrap.Hijack(w)
		if err != nil {
			log.Println(err)
			return
		}

		defer conn.Close()

		r.RequestURI = path

		reqbuf, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
			return
		}
		_ = reqbuf

		session := ag.NewSession()
		defer session.Close()
		chunkRequest := &api.ChunkRequest{
			Path:    path,
			Request: reqbuf,
		}
		sessionChunkerClient, err := session.Chunker(context.Background(), chunkRequest)
		if err != nil {
			log.Println(err)
			return
		}

		// TODO make a io.Reader from session.Chunker_Client, then call io.Copy
		for {
			chunk, err := sessionChunkerClient.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
				break
			}

			_, err = conn.Write(chunk.Chunk)
			if err != nil {
				log.Println(err)
				break
			}
		}
	}
}

func metricsRelay(ag types.Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			vars = mux.Vars(r)
			id   = vars["id"]
			path = strings.TrimPrefix(r.RequestURI, "/api/agent/"+id)
		)
		conn, err := wrap.Hijack(w)
		if err != nil {
			log.Println(err)
			return
		}

		defer conn.Close()

		r.RequestURI = path

		reqbuf, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
			return
		}
		_ = reqbuf

		session := ag.NewSession()
		defer session.Close()
		chunkRequest := &api.ChunkRequest{
			Path:    "metrics",
			Request: reqbuf,
		}
		sessionChunkerClient, err := session.Chunker(context.Background(), chunkRequest)
		if err != nil {
			log.Println(err)
			return
		}

		// TODO make a io.Reader from session.Chunker_Client, then call io.Copy
		for {
			chunk, err := sessionChunkerClient.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
				break
			}

			_, err = conn.Write(chunk.Chunk)
			if err != nil {
				log.Println(err)
				break
			}
		}
	}
}

func httpRelay(ag types.Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			vars      = mux.Vars(r)
			id        = vars["id"]
			portStr   = strings.TrimPrefix(r.RequestURI, "/api/agent/"+id+"/http/")
			port, err = strconv.Atoi(portStr)
		)
		if err != nil {
			http.Error(w, "port invalid: "+portStr, http.StatusBadRequest)
			return
		}
		_ = port
		// http.Error(w, "todo: tunneling data for you: "+portStr, http.StatusOK)
		conn, err := wrap.Hijack(w)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		session := ag.NewSession()
		defer session.Close()

		sessionForwardClient, err := session.Forward(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		// send addr info
		addrForwardMessage := &api.ForwardMessage{Head: ":" + portStr}
		err = sessionForwardClient.Send(addrForwardMessage)
		if err != nil {
			log.Println(err)
			return
		}

		// common error: ws transport is closing
		log.Println(forward(conn, sessionForwardClient))
	}
}

// (through chan Message{Type, Body} instead of interface)
func forward(ws io.ReadWriteCloser, session api.Session_ForwardClient) error {
	defer ws.Close()
	g, ctx := errgroup.WithContext(context.TODO())
	_ = ctx
	g.Go(func() error {
		log.Println("forward: client(ws) => session(grpc)")
		// TODO: io.Copy(session, ws), CopyBuffer, session.ReadFrom
		buf := make([]byte, 1<<12) // maximum input message is 4096 bytes
		for {
			n, err := ws.Read(buf)
			if err != nil {
				return err
			}
			msg := &api.ForwardMessage{Body: buf[:n]}
			err = session.Send(msg)
			if err != nil {
				return err
			}
		}
		return nil
	})

	g.Go(func() error {
		log.Println("forward: client(ws) <= session(grpc)")
		// TODO: io.Copy(ws, session), CopyBuffer, session.WriteTo
		for {
			resp, err := session.Recv()
			if err != nil {
				return err
			}
			_, err = ws.Write(resp.Body)
			if err != nil {
				return err
				break
			}
		}
		return nil
	})

	return g.Wait()
}
