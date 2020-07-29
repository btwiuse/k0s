package hub

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/btwiuse/pretty"
	"github.com/btwiuse/wetty/pkg/assets"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"k0s.io/k0s/pkg/api"
	"k0s.io/k0s/pkg/exporter"
    "k0s.io/k0s/pkg/wrap"
	types "k0s.io/k0s/pkg/hub"
	"modernc.org/httpfs"
	"nhooyr.io/websocket"
)

var (
	_ types.Hub = (*hub)(nil)
)

type hub struct {
	types.AgentManager

	*http.Server

	c              types.Config
	MetricsHandler http.Handler
}

func NewHub(c types.Config) types.Hub {
	var (
		listhand = NewHandleHijackListener()
		h        = &hub{
			c:              c,
			AgentManager:   NewAgentManager(),
			MetricsHandler: exporter.NewHandler(),
		}
	)
	// ensure core fields of h is not empty before return
	h.initServer(h.c.Port(), listhand)
	go h.serve(listhand, listhand)
	return h
}

func (h *hub) GetConfig() types.Config {
	return h.c
}

// this function is modeled after http.Serve(net.Listener, http.Handler)
// but unlike conventional servers, in which connections are extablished
// on the listener side and then passed on to handler,
// this one doesn't require listening on a port, and the direction in which
// connection goes is exactly opposite: the net.Conn's are created on the
// handler side and then sent through a (chan net.Conn) to the listener side
func (h *hub) serve(ln net.Listener, _ http.Handler) {
	// ln <- net.Conn <- hl
	// ln: conventionally a producer of net.Conn, but it's role here is consumer
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

	// unregister
	defer h.Del(rpc.ID())

	for {
		select {
		case f := <-rpc.Actions():
			go f(h)
		case <-time.After(3 * time.Second):
			go rpc.Ping()
		case <-rpc.Done():
			return
		}
	}
}

func (h *hub) initServer(addr string, hl http.Handler) {
	// http2 is not hijack friendly, use TLSNextProto to force HTTP/1.1
	h.Server = &http.Server{
		Addr:         addr,
		Handler:      h.initHandler(hl),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
}

func (h *hub) initHandler(hl http.Handler) http.Handler {
	r := mux.NewRouter()

	// list active agents
	r.HandleFunc("/api/agents/list", http.HandlerFunc(h.handleAgentsList)).Methods("GET")
	r.HandleFunc("/api/agents/watch", http.HandlerFunc(h.handleAgentsWatch)).Methods("GET")

	// client /api/agent/{id}/rootfs/{path} hijack => net.Conn <(copy) hijacked grpc fs conn
	// client /api/agent/{id}/ws => ws <(pipe)> hijacked grpc ws conn
	s := r.PathPrefix("/api/agent/{id}")
	s.Handler(http.HandlerFunc(h.handleAgent)).Methods("GET")

	// public api
	// agent hijack => yrpc -> hub.RPC -> hub.Agent
	// alternative websocket implementation:
	// http upgrade => websocket conn => net.Conn => hub.RPC -> hub.Agent

	// hl -> net.Conn -> ln
	// hl: conventionally a consumer of net.Conn, but it's role here is producer
	r.Handle("/api/rpc", hl).Methods("GET")

	// agent hijack => gRPC {ws, fs} -> hub.Session -> hub.Agent
	// alternative websocket implementation:
	// http upgrade => websocket conn => net.Conn => gRPC {ws, fs} -> hub.Session -> hub.Agent
	r.HandleFunc("/api/fs", h.handleTunnel(api.FS)).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/api/socks5", h.handleTunnel(api.Socks5)).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/api/redir", h.handleTunnel(api.Redir)).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/api/metrics", h.handleTunnel(api.Metrics)).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/api/terminal", h.handleTunnel(api.Terminal)).Methods("GET").Queries("id", "{id}")

	// hub specific function
	r.HandleFunc("/api/version", h.handleVersion).Methods("GET")
	r.Handle("/api/metrics", h.MetricsHandler).Methods("GET")
	return handlers.LoggingHandler(os.Stderr, cors.AllowAll().Handler(r))
}

func (h *hub) handleVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(pretty.JSONString(h.GetConfig().GetVersion())))
}

func (h *hub) handleAgentsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(pretty.JSONString(h.GetAgents())))
}

func (h *hub) handleAgentsWatch(w http.ResponseWriter, r *http.Request) {
	// conn, err := wrconn(w, r)
	wsconn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Println(err)
		return
	}
	conn := websocket.NetConn(context.Background(), wsconn, websocket.MessageText)

	watchInterval := time.Second
	for {
		_, err := conn.Write([]byte(pretty.JSONString(h.GetAgents())))
		if err != nil {
			log.Println("agents watch:", err)
			break
		}
		time.Sleep(watchInterval)
	}
}

func (h *hub) handleTunnel(tun api.Tunnel) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			vars = mux.Vars(r)
			id   = vars["id"]
		)

		if !h.Has(id) {
			log.Println("no such id", id)
			return
		}

		conn, err := wrap.Wrconn(w, r)
		if err != nil {
			log.Printf("error accepting %s: %s\n", tun, err)
			return
		}

		h.GetAgent(id).AddTunnel(tun, conn)
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
		http.Redirect(w, r, "/", http.StatusFound /*302*/)
		return
	}

	ag := h.GetAgent(id)

	// delegate := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
	switch {
	case strings.HasPrefix(subpath, "/rootfs"):
		ag.BasicAuth(http.HandlerFunc(fsRelay(ag))).ServeHTTP(w, r)
	case strings.HasPrefix(subpath, "/redir"):
		ag.BasicAuth(http.HandlerFunc(redirRelay(ag))).ServeHTTP(w, r)
	case strings.HasPrefix(subpath, "/socks5"):
		ag.BasicAuth(http.HandlerFunc(socks5Relay(ag))).ServeHTTP(w, r)
	case strings.HasPrefix(subpath, "/metrics"):
		metricsRelay(ag)(w, r)
	case strings.HasPrefix(subpath, "/terminal"):
		ag.BasicAuth(http.HandlerFunc(terminalRelay(ag))).ServeHTTP(w, r)
	default:
		ag.BasicAuth(staticFileHandler).ServeHTTP(w, r)
	}
}
