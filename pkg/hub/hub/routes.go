package hub

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/btwiuse/conntroll/pkg/api"
	types "github.com/btwiuse/conntroll/pkg/hub"
	"github.com/btwiuse/conntroll/pkg/hub/agent"
	"github.com/btwiuse/conntroll/pkg/wrap"
	"github.com/btwiuse/pretty"
	"github.com/btwiuse/wetty/pkg/assets"
	"github.com/btwiuse/wetty/pkg/msg"
	"github.com/btwiuse/wetty/pkg/utils"
	"github.com/btwiuse/wetty/pkg/wetty"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
	"modernc.org/httpfs"
)

var (
	_ types.Hub = (*hub)(nil)
)

type hub struct {
	types.AgentManager

	*http.Server
}

func NewHub(c types.Config) types.Hub {
	h := &hub{
		AgentManager: NewAgentManager(),
	}
	h.serve(c.Port())
	return h
}

// TODO better gorilla mux router
func (h *hub) serve(addr string) {
	r := mux.NewRouter()
	r.HandleFunc("/api/agents/", h.handleAgents).Methods("GET")
	s := r.PathPrefix("/api/agent/{id}")
	s.HandlerFunc(h.handleAgent).Methods("GET")
	r.HandleFunc("/api/rpc", h.handleRPC).Methods("GET").
		Queries(
			"id", "{id}",
			"pwd", "{pwd}",
			"os", "{os}",
			"arch", "{arch}",
			"username", "{username}",
			"hostname", "{hostname}",
		)
	r.HandleFunc("/api/session", h.handleSession).Methods("GET").
		Queries(
			"id", "{id}",
		)
	h.Server = &http.Server{
		Addr:         addr,
		Handler:      handlers.LoggingHandler(os.Stderr, cors.Default().Handler(r)),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
}

func (h *hub) handleAgents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(pretty.JSONString(h.GetAgents())))
}

func (h *hub) handleAgent(w http.ResponseWriter, r *http.Request) {
	var (
		vars                           = mux.Vars(r)
		id                             = vars["id"]
		subpath                        = strings.TrimPrefix(r.RequestURI, "/api/agent/"+id)
		staticFileServer  http.Handler = http.FileServer(httpfs.NewFileSystem(assets.Assets, time.Now()))
		staticFileHandler http.Handler = http.StripPrefix("/api/agent/"+id+"/", staticFileServer)
	)

	log.Println("handleAgent", id, subpath)

	if !h.Has(id) {
		log.Println("hub has no such id", id)
		for i, ider := range h.Values() {
			log.Println(i, ider.ID())
		}
		http.Redirect(w, r, "/", 302)
		return
	}

	agent := h.GetAgent(id)

	switch {
	case strings.HasPrefix(subpath, "/ws"):
		wsRelay(agent)(w, r)
	case strings.HasPrefix(subpath, "/rootfs"):
		fsRelay(agent)(w, r)
	default:
		staticFileHandler.ServeHTTP(w, r)
	}
}

func wsRelay(ag types.Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wsconn, err := wetty.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer wsconn.Close()

		session := ag.NewSession()
		sessionSendClient, err := session.Send(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(pipe(utils.WsConnToReadWriter(wsconn), sessionSendClient))
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

		conn, err := wrap.HijackConn(w.(http.Hijacker).Hijack())
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

		session := ag.NewSession()
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

func (h *hub) handleRPC(w http.ResponseWriter, r *http.Request) {
	log.Println("handleRPC")
	conn, err := wrap.HijackConn(w.(http.Hijacker).Hijack())
	if err != nil {
		log.Println("error hijacking:", err)
		return
	}

	var (
		vars     = mux.Vars(r)
		id       = vars["id"]
		pwd      = vars["pwd"]
		ip, _, _ = net.SplitHostPort(conn.RemoteAddr().String())
		username = vars["username"]
		hostname = vars["hostname"]
		goos     = vars["os"]
		goarch   = vars["arch"]
	)

	if h.Has(id) {
		h.GetAgent(id).AddRPCConn(conn)
		return
	}

	ag := agent.NewAgent(conn,
		agent.SetID(id),
		agent.SetIP(ip),
		agent.SetPWD(pwd),
		agent.SetUsername(username),
		agent.SetHostname(hostname),
		agent.SetOS(goos),
		agent.SetARCH(goarch),
	)
	h.Add(ag)
	go h.GC(ag)
}

func (h *hub) GC(ag types.Agent) {
	select {
	case <-ag.Done():
		h.Del(ag.ID())
	}
}

func (h *hub) handleSession(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		id   = vars["id"]
	)

	conn, err := wrap.HijackConn(w.(http.Hijacker).Hijack())
	if err != nil {
		log.Println("error hijacking:", err)
		return
	}

	if !h.Has(id) {
		log.Println("no such id", id)
		return
	}

	h.GetAgent(id).AddSessionConn(conn)
}
