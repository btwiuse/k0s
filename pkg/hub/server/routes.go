package server

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
	"time"

	"github.com/btwiuse/pretty"
	"github.com/btwiuse/sse"
	"github.com/btwiuse/wsconn"
	"github.com/gorilla/mux"
	echo "github.com/jpillora/go-echo-server/handler"
	"k0s.io"
	"k0s.io/pkg/api"
	"k0s.io/pkg/hub"
	"k0s.io/pkg/log"
	"k0s.io/pkg/middleware"
	"k0s.io/pkg/ui"
)

var (
	_ hub.Hub = (*hubServer)(nil)
)

type hubServer struct {
	hub.AgentManager

	*http.Server

	c          hub.Config
	BinHandler http.Handler
	sseMux     *sse.SSE
}

func (h *hubServer) Handler() http.Handler {
	return h.Server.Handler
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

func NewHub(c hub.Config) hub.Hub {
	var (
		listhand = NewHTTPChannelListener()
		h        = &hubServer{
			c:            c,
			AgentManager: NewAgentManager(),
			BinHandler:   middleware.GzipMiddleware(binHandler()),
			sseMux:       sse.NewSSE(),
		}
	)
	go func() {
		for {
			h.sseMux.SetData(pretty.JSONStringLine(h.GetAgents()))
			time.Sleep(time.Second)
		}
	}()
	// ensure core fields of h is not empty before return
	h.setupServer(h.GetConfig().Port(), "/api", listhand)
	go h.serveLoop(listhand)
	return h
}

func (h *hubServer) GetConfig() hub.Config {
	return h.c
}

// this function is modeled after http.Serve(net.Listener, http.Handler)
// but unlike conventional servers, in which connections are extablished
// on the listener side and then passed on to handler,
// this one doesn't require listening on a port, and the direction in which
// connection goes is exactly opposite: the net.Conn's are created on the
// handler side and then sent through a (chan net.Conn) to the listener side
func (h *hubServer) serveLoop(ln net.Listener) {
	// ln <- net.Conn <- hl
	// ln: conventionally a producer of net.Conn, but it's role here is consumer
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go h.upgrade(conn)
	}
}

func (h *hubServer) upgrade(conn net.Conn) {
	var ss = NewServerSession(conn)

	// unregister
	defer h.Del(ss.ID())

	for {
		select {
		case f := <-ss.Actions():
			go f(h)
		case <-time.After(3 * time.Second):
			go ss.Ping()
		case <-ss.Done():
			return
		}
	}
}

func (h *hubServer) setupServer(addr, apiPrefix string, hl http.Handler) {
	handler := middleware.AllowAllCorsMiddleware(h.installRoutes(apiPrefix, hl))
	if h.GetConfig().Verbose() {
		handler = middleware.LoggingMiddleware(handler)
	}
	// http2 is not hijack friendly, use TLSNextProto to force HTTP/1.1
	h.Server = &http.Server{
		Addr:         addr,
		Handler:      handler,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
}

func (h *hubServer) installRoutes(apiPrefix string, hl http.Handler) (R *mux.Router) {
	if h.GetConfig().UI() {
		R = ui.NewRouter(k0s.DEFAULT_UI_ADDRESS)
	} else {
		R = mux.NewRouter()
	}

	r := R.PathPrefix(apiPrefix).Subrouter()

	r.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	r.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	r.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	r.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	r.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

	// list active agents
	r.Handle("/agents/list", http.HandlerFunc(h.handleAgentsList)).Methods("GET")
	r.Handle("/agents/watch", h.sseMux).Methods("GET")

	// client /api/agent/{id}/rootfs/{path} hijack => net.Conn <(copy) hijacked grpc fs conn
	// client /api/agent/{id}/ws => ws <(pipe)> hijacked grpc ws conn
	s := r.PathPrefix("/agent/{id}")
	s.Handler(http.HandlerFunc(h.handleAgent)) // allow all methods

	// order routes from most specific to least specific
	r.HandleFunc("/upgrade", h.handleStreamUpgrade).Methods("GET").Queries("id", "{id}").Queries("protocol", "{protocol}")
	r.Handle("/upgrade", hl).Methods("GET")

	// dev helper
	r.Handle("/echo", echo.New(echo.Config{})).Methods(
		http.MethodGet,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodPost,
	)

	// hub specific function
	r.HandleFunc("/version", h.handleVersion).Methods("GET")
	r.Handle("/bin/k0s", h.BinHandler).Methods("GET")

	return R
}

func (h *hubServer) handleVersion(w http.ResponseWriter, r *http.Request) {
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

func (h *hubServer) handleAgentsList(w http.ResponseWriter, r *http.Request) {
	var (
		// vars = mux.Vars(r)
		// tags = vars["tags"]
		vars        = r.URL.Query()
		_, untagged = vars["untagged"]
		vtags       = vars.Get("tags")
		tags        = strings.Split(vtags, ",")
		allAgents   = h.GetAgents()
		agents      = []hub.Agent{}
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

func (h *hubServer) handleStreamUpgrade(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		p    = api.ProtocolID(vars["protocol"])
		id   = vars["id"]
	)
	println("handleStreamUpgrade", string(p))

	if !h.Has(id) {
		log.Println("no such id", id)
		return
	}
	ag := h.GetAgent(id)

	conn, err := wsconn.Wrconn(w, r)
	if err != nil {
		log.Printf("error accepting %s: %s\n", p, err)
		return
	}

	ag.ChannelChan(p) <- conn
}

func (h *hubServer) handleAgent(w http.ResponseWriter, r *http.Request) {
	var (
		vars    = mux.Vars(r)
		id      = vars["id"]
		subpath = strings.TrimPrefix(r.RequestURI, "/api/agent/"+id)
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
	case strings.HasPrefix(subpath, "/fsv2"):
		ag.BasicAuth(http.HandlerFunc(fsV2Relay(ag))).ServeHTTP(w, r)
	case strings.HasPrefix(subpath, "/rootfs"):
		ag.BasicAuth(http.HandlerFunc(fsRelay(ag))).ServeHTTP(w, r)
	case strings.HasPrefix(subpath, "/redir"):
		ag.BasicAuth(http.HandlerFunc(redirRelay(ag))).ServeHTTP(w, r)
	case strings.HasPrefix(subpath, "/socks5"):
		ag.BasicAuth(http.HandlerFunc(socks5Relay(ag))).ServeHTTP(w, r)
	case strings.HasPrefix(subpath, "/version"):
		versionRelay(ag)(w, r)
	case strings.HasPrefix(subpath, "/env"):
		envRelay(ag)(w, r)
	case strings.HasPrefix(subpath, "/doh"):
		dohRelay(ag)(w, r)
	case strings.HasPrefix(subpath, "/jsonl"):
		jsonlRelay(ag)(w, r)
	case strings.HasPrefix(subpath, "/xpra"):
		xpraRelay(ag)(w, r)
	case strings.HasPrefix(subpath, "/terminalv2"): // must come before "/terminal" otherwise won't ever match
		ag.BasicAuth(http.HandlerFunc(terminalV2Relay(ag))).ServeHTTP(w, r)
	case strings.HasPrefix(subpath, "/terminal"):
		ag.BasicAuth(http.HandlerFunc(terminalRelay(ag))).ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}
