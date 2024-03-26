package handler

import (
	"bufio"
	"context"
	"expvar"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"k0s.io/pkg/log"
	"k0s.io/pkg/middleware"
	portless "k0s.io/pkg/tunnel"
	"github.com/btwiuse/wsconn"
)

func Handler(prefix string) http.Handler {
	r := http.NewServeMux()
	r.Handle("/", manageHub(setupTunnel(defaultTunnelMux)))
	return middleware.LoggingMiddleware(middleware.AllowAllCorsMiddleware(r))
}

func paths() []string {
	st := []string{}
	defaultTunnelMux.mu.RLock()
	for k := range defaultTunnelMux.TunnelConns {
		st = append(st, k)
	}
	defaultTunnelMux.mu.RUnlock()
	return st
}

func keys() []string {
	st := []string{}
	defaultTunnelMux.mu.RLock()
	for k := range defaultTunnelMux.Conns {
		st = append(st, k)
	}
	defaultTunnelMux.mu.RUnlock()
	return st
}

func manageHub(next http.Handler) http.Handler {
	expvar.Publish("Number of goroutines", expvar.Func(func() interface{} { return runtime.NumGoroutine() }))
	expvar.Publish("Registered ids", expvar.Func(func() interface{} { return keys() }))
	expvar.Publish("Registered routes", expvar.Func(func() interface{} { return paths() }))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := http.NewServeMux()
		m.Handle("/", next)
		m.Handle("/_/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: use html template to render the page
			io.WriteString(w, "<h1>Management Console</h1>\n")
			io.WriteString(w, "<h2>Registered Mounts</h2>\n")
			defaultTunnelMux.mu.RLock()
			entries := []string{}
			for p, e := range defaultTunnelMux.m {
				entries = append(entries, fmt.Sprintf("<p><a href=\"%s\">%s</a> => %s</p>\n", p, p, e.from))
			}
			// TODO: sort entries instead of templated strings
			sort.Strings(entries)
			for _, entry := range entries {
				io.WriteString(w, entry)
			}
			defaultTunnelMux.mu.RUnlock()
			io.WriteString(w, "<h2>Debug Info</h2>\n")
			io.WriteString(w, fmt.Sprintf("<p><a href=\"%s\">%s</a></p>\n", "/_/", "/_/"))
		}))
		m.Handle("/_/expvar", expvar.Handler())
		m.ServeHTTP(w, r)
	})
}

// get stream identifier from query string if specified
// otherwise from request header
func getFingerprint(r *http.Request) string {
	param := r.URL.Query().Get(portless.FingerPrintHeader)
	if param != "" {
		return param
	}
	return r.Header.Get(portless.FingerPrintHeader)
}

// split internal tunnel requests from normal requests
// register custom routes
func setupTunnel(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fp := getFingerprint(r)
		log.Println("fingerprint", fp)
		pattern := r.URL.Path
		from := r.URL.Query().Get("from")
		switch _, ok := defaultTunnelMux.Tunnels[fp]; {
		// passthrough regular request to wsconnped handler
		case fp == "":
			next.ServeHTTP(w, r)
		// new stream identifier, create new tunnel
		case !ok:
			defaultTunnelMux.Tunnels[fp] = portless.NewTunnel()
			// use first conn as controlling channel
			go func() {
				conn, err := defaultTunnelMux.Tunnels[fp].Accept()
				if err != nil {
					log.Println(err)
					return
				}
				defaultTunnelMux.Conns[fp] = conn
				defaultTunnelMux.TunnelConns[pattern] = conn
				defaultTunnelMux.Handle(pattern, from, defaultTunnelMux.Tunnels[fp])
				go gc(fp, pattern, conn)
			}()
			log.Printf("[SRV:%s] %s => %s\n", fp, r.RemoteAddr+"@"+from, r.URL.Path)
			fallthrough
		// handle request with registered
		default:
			defaultTunnelMux.Tunnels[fp].ServeHTTP(w, r)
		}
	})
}

func gc(id string, pattern string, conn net.Conn) {
	defer func() {
		log.Println("hub: srv lost", id, pattern)
		defaultTunnelMux.Delete(id, pattern)
	}()
	for bufio.NewScanner(conn).Scan() {
		// log.Println("ping")
	}
}

var defaultTunnelMux = &tunnelMux{
	mu: &sync.RWMutex{},
	m:  make(map[string]muxEntry),

	Conns:       make(map[string]net.Conn),
	TunnelConns: make(map[string]net.Conn),
	Tunnels:     make(map[string]*portless.Tunnel),
}

// modeled after http.ServeMux
type tunnelMux struct {
	mu *sync.RWMutex
	m  map[string]muxEntry
	es []muxEntry // slice of entries sorted from longest to shortest

	Conns       map[string]net.Conn
	TunnelConns map[string]net.Conn
	Tunnels     map[string]*portless.Tunnel
}

// Find a handler on a handler map given a path string.
// Most-specific (longest) pattern wins.
func (mux *tunnelMux) match(path string) *muxEntry {
	mux.mu.RLock()
	defer mux.mu.RUnlock()
	// Check for exact match first.
	v, ok := mux.m[path]
	if ok {
		return &v
	}

	// Check for longest valid match.  mux.es contains all patterns
	// that end in / sorted from longest to shortest
	for _, e := range mux.es {
		if strings.HasPrefix(path, e.pattern) {
			return &e
		}
	}
	return nil
}

// ServeHTTP dispatches the request to the handler whose
// pattern most closely matches the request URL.
func (mux *tunnelMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	isWS := strings.Split(r.Header.Get("Upgrade"), ",")[0] == "websocket"

	path := r.URL.Path
	e := defaultTunnelMux.match(path)
	if e == nil {
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}

	_, err := io.WriteString(defaultTunnelMux.TunnelConns[e.pattern], "ACCEPT\n")
	if err != nil {
		log.Println(err)
		return
	}

	// here we capture the http request,
	// or use http.Request.Write
	// Google: how caddy reverse proxy works?

	if isWS {
		connSrc, err := e.tunnel.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		connDst, err := wsconn.Wrconn(w, r)
		if err != nil {
			log.Println(err)
			return
		}

		go io.Copy(connSrc, connDst)
		go io.Copy(connDst, connSrc)
		return
	}

	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return e.tunnel.Accept()
		},
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	rp := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			// modify request Path
			log.Println("director: request url:", r.URL, r.Host)
			req.URL.Scheme = "http" //r.URL.Scheme
			req.Host = r.Host
			req.URL.Host = req.Host
			// req.URL.Path = rebase(req.URL.Path, e.pattern, e.from)
			// req.RequestURI = rebase(req.RequestURI, e.pattern, e.from)
		},
		Transport: tr,
	}

	rp.ServeHTTP(w, r)
}

// modeled after
type muxEntry struct {
	tunnel  *portless.Tunnel
	pattern string
	from    string
}

// Handle registers the handler for the given pattern.
// If a handler already exists for pattern, Handle panics.
func (mux *tunnelMux) Handle(pattern, from string, tunnel *portless.Tunnel) {
	mux.mu.Lock()
	mux.m[pattern] = muxEntry{tunnel: tunnel, pattern: pattern, from: from}
	mux.mu.Unlock()
	if pattern[len(pattern)-1] == '/' {
		mux.updateEntries()
	}
}

func (mux *tunnelMux) Delete(id, pattern string) {
	mux.mu.Lock()
	delete(mux.Conns, id)
	delete(mux.Tunnels, id)
	delete(mux.TunnelConns, pattern)
	delete(mux.m, pattern)
	mux.mu.Unlock()
	mux.updateEntries()
}

func (mux *tunnelMux) updateEntries() {
	mux.mu.Lock()
	defer mux.mu.Unlock()
	es := []muxEntry{}
	for pattern, e := range mux.m {
		if pattern[len(pattern)-1] == '/' {
			es = appendSorted(es, e)
		}
	}
	mux.es = es
}

func appendSorted(es []muxEntry, e muxEntry) []muxEntry {
	n := len(es)
	i := sort.Search(n, func(i int) bool {
		return len(es[i].pattern) < len(e.pattern)
	})
	if i == n {
		return append(es, e)
	}
	// we know that i points at where we want to insert
	es = append(es, muxEntry{}) // try to grow the slice in place, any entry works.
	copy(es[i+1:], es[i:])      // Move shorter entries down
	es[i] = e
	return es
}

func rebase(path, pattern, from string) string {
	// initialize variable
	newpath := path

	fromURL, err := url.Parse(from)
	if err != nil {
		return newpath
	}

	// path begins with pattern
	if pattern[len(pattern)-1] == '/' && from[len(from)-1] == '/' {
		newpath = fromURL.Path + strings.TrimPrefix(path, pattern)
	}

	// path == pattern
	if pattern[len(pattern)-1] != '/' && from[len(from)-1] != '/' {
		newpath = fromURL.Path
	}

	log.Println(fmt.Sprintf("rewrite path %s => %s\n", path, newpath))

	return newpath
}
