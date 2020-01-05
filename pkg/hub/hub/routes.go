package hub

import (
	"context"
	"crypto/subtle"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	types "github.com/btwiuse/conntroll/pkg/hub"
	"github.com/btwiuse/pretty"
	"github.com/btwiuse/wetty/pkg/assets"
	webui "github.com/conntroll/conntroll.github.io"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"modernc.org/httpfs"
	"nhooyr.io/websocket"
)

var (
	_ types.Hub = (*hub)(nil)
)

type hub struct {
	types.AgentManager

	*http.Server

	ba        bool
	localui   bool // load webui assets from local dir for debugging purpose
	user      string
	pass      string
	handleRPC http.Handler // http.Handler|net.Listener
}

func NewHub(c types.Config) types.Hub {
	h := &hub{
		AgentManager: NewAgentManager(),
		localui:      c.LocalUI(),
	}
	h.startRPCServer()
	h.user, h.pass, h.ba = c.BasicAuth()
	h.initServer(c.Port())
	return h
}

// https://stackoverflow.com/questions/21936332/idiomatic-way-of-requiring-http-basic-auth-in-go/39591234#39591234
func (h *hub) basicauth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if h.ba {
			username, password, ok := r.BasicAuth()
			log.Println("basicauth:", username, password, ok)
			if !ok || subtle.ConstantTimeCompare([]byte(h.user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(h.pass), []byte(password)) != 1 {
				realm := "please enter hub password"
				w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
				w.WriteHeader(401)
				w.Write([]byte("Unauthorised.\n"))
				return
			}
		}
		next.ServeHTTP(w, r)
	}
}

func (h *hub) serveRPC(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go h.register(conn)
	}
}

func (h *hub) register(conn net.Conn) {
	var rpc = ToRPC(conn)

	defer rpc.Unregister(h)

	for {
		select {
		case f := <-rpc.Actions():
			go f(h)
		case <-rpc.Done():
			return
		case <-time.After(3 * time.Second):
			go rpc.Ping()
		}
	}
}

func (h *hub) initServer(addr string) {
	r := mux.NewRouter()

	if h.localui {
		r.NotFoundHandler = h.basicauth(http.FileServer(http.Dir("conntroll.github.io")))
	} else {
		mergedAssets := assets.Assets
		for k, v := range webui.Assets {
			mergedAssets[k] = v
		}
		mergedAssets["favicon.ico"] = ""
		r.NotFoundHandler = h.basicauth(http.FileServer(httpfs.NewFileSystem(mergedAssets, time.Now())))
	}

	// list active agents
	r.HandleFunc("/api/agents/list", h.basicauth(http.HandlerFunc(h.handleAgentsList))).Methods("GET")
	r.HandleFunc("/api/agents/watch", h.basicauth(http.HandlerFunc(h.handleAgentsWatch))).Methods("GET")

	// client /api/agent/{id}/rootfs/{path} hijack => net.Conn <(copy) hijacked grpc fs conn
	// client /api/agent/{id}/ws => ws <(pipe)> hijacked grpc ws conn
	s := r.PathPrefix("/api/agent/{id}")
	// s.Handler(h.basicauth(http.HandlerFunc(h.handleAgent))).Methods("GET")
	s.Handler(http.HandlerFunc(h.handleAgent)).Methods("GET")

	// public api
	// agent hijack => yrpc -> hub.RPC -> hub.Agent
	// alternative websocket implementation:
	// http upgrade => websocket conn => net.Conn => hub.RPC -> hub.Agent
	r.Handle("/api/yrpc", h.handleRPC).Methods("GET") // TODO: remove this legacy endpoint in the future
	r.Handle("/api/rpc", h.handleRPC).Methods("GET")

	// agent hijack => gRPC {ws, fs} -> hub.Session -> hub.Agent
	// alternative websocket implementation:
	// http upgrade => websocket conn => net.Conn => gRPC {ws, fs} -> hub.Session -> hub.Agent
	r.HandleFunc("/api/grpc", h.handleGRPC).Methods("GET").Queries("id", "{id}")

	// http2 is not hijack friendly, use TLSNextProto to force HTTP/1.1
	h.Server = &http.Server{
		Addr:         addr,
		Handler:      handlers.LoggingHandler(os.Stderr, cors.AllowAll().Handler(r)),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
}

func (h *hub) handleAgentsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(pretty.JSONString(h.GetAgents())))
}

func (h *hub) handleAgentsWatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	wsconn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})

	if err != nil {
		log.Println(err)
		return
	}

	conn := websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary)

	for range time.Tick(time.Second) {
		_, err := conn.Write([]byte(pretty.JSONString(h.GetAgents())))
		if err != nil {
			log.Println("agents watch:", err)
			break
		}
	}
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

	ag := h.GetAgent(id)

	// delegate := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
	switch {
	case strings.HasPrefix(subpath, "/ws"):
		ag.BasicAuth(http.HandlerFunc(wsRelay(ag))).ServeHTTP(w, r)
	case strings.HasPrefix(subpath, "/rootfs"):
		ag.BasicAuth(http.HandlerFunc(fsRelay(ag))).ServeHTTP(w, r)
	case strings.HasPrefix(subpath, "/metrics"):
		metricsRelay(ag)(w, r)
	case h.localui:
		ag.BasicAuth(staticFileHandler).ServeHTTP(w, r)
	default:
		ag.BasicAuth(agentIndexHandler(ag)).ServeHTTP(w, r)
	}
}

func agentIndexHandler(ag types.Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf(`<!doctype html>
<html>
  <head>
    <title>%s@%s</title>
    <link rel="icon" type="image/x-icon" href="/favicon.ico">
    <link rel="stylesheet" href="/css/index.css" />
    <link rel="stylesheet" href="/css/xterm.css" />
    <link rel="stylesheet" href="/css/xterm_customize.css" />
  </head>
  <body>
    <div id="terminal"></div>
    <script src="/js/wetty-bundle.js"></script>
  </body>
</html>
`, ag.GetUsername(), ag.GetHostname())))
	}
}

func (h *hub) handleGRPC(w http.ResponseWriter, r *http.Request) {
	var (
		vars   = mux.Vars(r)
		id     = vars["id"]
		wsconn *websocket.Conn
		conn   net.Conn
		err    error
	)

	if !h.Has(id) {
		log.Println("no such id", id)
		return
	}

	wsconn, err = websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	conn = websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary)

	if err != nil {
		log.Println("error accepting grpc:", err)
		return
	}

	h.GetAgent(id).AddSessionConn(conn)
}

func (h *hub) startRPCServer() {
	var (
		listhand              = NewHandleHijackListener()
		handler  http.Handler = listhand
		listener net.Listener = listhand
	)
	h.handleRPC = handler
	go h.serveRPC(listener)
}
