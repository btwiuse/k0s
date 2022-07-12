package hub

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
	"time"

	"github.com/btwiuse/pretty"
	"github.com/btwiuse/wetty/pkg/assets"
	"github.com/gorilla/mux"
	echo "github.com/jpillora/go-echo-server/handler"
	"k0s.io/pkg/api"
	"k0s.io/pkg/exporter"
	types "k0s.io/pkg/hub"
	"k0s.io/pkg/middleware"
	"k0s.io/pkg/wrap"
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
	BinHandler     http.Handler
}

func binHandler() http.Handler {
	exe, err := os.Executable()
	if err != nil {
		return http.NotFoundHandler()
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, exe)
	})
}

func NewHub(c types.Config) types.Hub {
	var (
		listhand = NewHandleHijackListener()
		h        = &hub{
			c:              c,
			AgentManager:   NewAgentManager(),
			MetricsHandler: middleware.GzipMiddleware(exporter.NewHandler()),
			BinHandler:     middleware.GzipMiddleware(binHandler()),
		}
	)
	// ensure core fields of h is not empty before return
	h.initServer(h.c.Port(), "/api", listhand)
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

func (h *hub) initServer(addr, apiPrefix string, hl http.Handler) {
	// http2 is not hijack friendly, use TLSNextProto to force HTTP/1.1
	h.Server = &http.Server{
		Addr:         addr,
		Handler:      h.initHandler(apiPrefix, hl),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
}

func (h *hub) initHandler(apiPrefix string, hl http.Handler) http.Handler {
	r := mux.NewRouter().PathPrefix(apiPrefix).Subrouter()

	r.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	r.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	r.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	r.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	r.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

	// list active agents
	r.Handle("/agents/list", http.HandlerFunc(h.handleAgentsList)).Methods("GET")
	r.Handle("/agents/watch", http.HandlerFunc(h.handleAgentsWatch)).Methods("GET")

	// client /api/agent/{id}/rootfs/{path} hijack => net.Conn <(copy) hijacked grpc fs conn
	// client /api/agent/{id}/ws => ws <(pipe)> hijacked grpc ws conn
	s := r.PathPrefix("/agent/{id}")
	s.Handler(http.HandlerFunc(h.handleAgent)) // allow all methods

	// public api
	// agent hijack => yrpc -> hub.RPC -> hub.Agent
	// alternative websocket implementation:
	// http upgrade => websocket conn => net.Conn => hub.RPC -> hub.Agent

	// hl -> net.Conn -> ln
	// hl: conventionally a consumer of net.Conn, but it's role here is producer
	r.Handle("/rpc", hl).Methods("GET")

	// dev helper
	r.Handle("/echo", echo.New(echo.Config{})).Methods(
		http.MethodGet,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodPost,
	)

	// agent hijack => gRPC {ws, fs} -> hub.Session -> hub.Agent
	// alternative websocket implementation:
	// http upgrade => websocket conn => net.Conn => gRPC {ws, fs} -> hub.Session -> hub.Agent
	r.HandleFunc("/fs", h.handleTunnel(api.FS)).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/socks5", h.handleTunnel(api.Socks5)).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/redir", h.handleTunnel(api.Redir)).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/metrics", h.handleTunnel(api.Metrics)).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/k16s", h.handleTunnel(api.K16s)).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/doh", h.handleTunnel(api.Doh)).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/env", h.handleTunnel(api.Env)).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/terminal", h.handleTunnel(api.Terminal)).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/terminalv2", h.handleTunnel(api.TerminalV2)).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/version", h.handleTunnel(api.Version)).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/jsonl", h.handleTunnel(api.Jsonl)).Methods("GET").Queries("id", "{id}")

	// hub specific function
	r.HandleFunc("/version", h.handleVersion).Methods("GET")
	r.Handle("/metrics", h.MetricsHandler).Methods("GET")
	r.Handle("/bin/k0s", h.BinHandler).Methods("GET")
	// return middleware.LoggingMiddleware(middleware.AllowAllCorsMiddleware(r))
	return middleware.AllowAllCorsMiddleware(r)
}

func (h *hub) handleVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(pretty.JSONStringLine(h.GetConfig().GetVersion())))
}

func contains(set []string, e string) bool {
	for _, s := range set {
		if s == e {
			return true
		}
	}
	return false
}

func containsAll(set []string, subset []string) bool {
	for _, se := range subset {
		if !contains(set, se) {
			return false
		}
	}
	return true
}

func (h *hub) handleAgentsList(w http.ResponseWriter, r *http.Request) {
	var (
		// vars = mux.Vars(r)
		// tags = vars["tags"]
		vars        = r.URL.Query()
		_, untagged = vars["untagged"]
		vtags       = vars.Get("tags")
		tags        = strings.Split(vtags, ",")
		allAgents   = h.GetAgents()
		agents      = []types.Agent{}
	)
	switch {
	case untagged:
		for _, a := range allAgents {
			if len(a.GetTags()) == 0 {
				agents = append(agents, a)
			}
		}
	case vtags != "":
		for _, a := range allAgents {
			if containsAll(a.GetTags(), tags) {
				agents = append(agents, a)
			}
		}
	default:
		agents = allAgents
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(pretty.JSONStringLine(agents)))
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
		_, err := conn.Write([]byte(pretty.JSONStringLine(h.GetAgents())))
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

	// TODO: lookup agent by name
	if !h.Has(id) {
		//  log.Println("hub has no such id", id)
		//  for i, ider := range h.Values() {
		//  	log.Println(i, ider.ID())
		//  }
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
	case strings.HasPrefix(subpath, "/version"):
		versionRelay(ag)(w, r)
	case strings.HasPrefix(subpath, "/k16s"):
		k16sRelay(ag)(w, r)
	case strings.HasPrefix(subpath, "/env"):
		envRelay(ag)(w, r)
	case strings.HasPrefix(subpath, "/doh"):
		dohRelay(ag)(w, r)
	case strings.HasPrefix(subpath, "/jsonl"):
		jsonlRelay(ag)(w, r)
	case strings.HasPrefix(subpath, "/metrics"):
		var (
			vars    = r.URL.Query()
			_, k16s = vars["k16s"]
		)
		if k16s {
			k16sRelay(ag)(w, r)
		} else {
			metricsRelay(ag)(w, r)
		}
	case strings.HasPrefix(subpath, "/terminalv2"): // must come before "/terminal" otherwise won't ever match
		ag.BasicAuth(http.HandlerFunc(terminalV2Relay(ag))).ServeHTTP(w, r)
	case strings.HasPrefix(subpath, "/terminal"):
		ag.BasicAuth(http.HandlerFunc(terminalRelay(ag))).ServeHTTP(w, r)
	default:
		ag.BasicAuth(staticFileHandler).ServeHTTP(w, r)
	}
}
