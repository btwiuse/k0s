package hub

import (
	"bufio"
	"context"
	"crypto/subtle"
	"crypto/tls"
	"fmt"
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
	agentinfo "github.com/btwiuse/conntroll/pkg/hub/agent/info"
	"github.com/btwiuse/conntroll/pkg/rng"
	"github.com/btwiuse/conntroll/pkg/uuid"
	"github.com/btwiuse/conntroll/pkg/wrap"
	"github.com/btwiuse/pretty"
	"github.com/btwiuse/wetty/pkg/assets"
	"github.com/btwiuse/wetty/pkg/msg"
	"github.com/btwiuse/wetty/pkg/wetty"
	webui "github.com/conntroll/conntroll.github.io"
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

	ba         bool
	localui    bool // load webui assets from local dir for debugging purpose
	user       string
	pass       string
	handleYRPC http.Handler // http.Handler|net.Listener
}

func NewHub(c types.Config) types.Hub {
	h := &hub{
		AgentManager: NewAgentManager(),
		localui:      c.LocalUI(),
	}
	h.startYRPCServer()
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

func (h *hub) serveYRPC(ln net.Listener) {
	for {
		log.Println("listening YRPC conns")
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// parse agent info
		scanner := bufio.NewScanner(conn)
		if !scanner.Scan() {
			continue
		}

		buf := scanner.Bytes()
		info, err := agentinfo.Decode(buf)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println(info)

		go h.toAgent(info, conn)
	}
}

func (h *hub) toAgent(info types.Info, conn net.Conn) {
	if h.Has(info.GetID()) {
		io.WriteString(conn, "duplicate id\n")
		return
	}

	var (
		rpc = ToRPC(conn)
		ag  = agent.NewAgent(rpc, info)
	)

	h.Add(ag)

	defer func() {
		h.Del(info.GetID())
		ag.Close()
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		log.Println(info.GetID(), scanner.Text())
	}
}

func (h *hub) initServer(addr string) {
	r := mux.NewRouter()

	// ==================== basic auth (TODO) =======================
	// root auth

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

	// ========================== public ============================
	// agent hijack => yrpc -> hub.RPC -> hub.Agent
	r.Handle("/api/yrpc", h.handleYRPC).Methods("GET")

	// agent hijack => gRPC {ws, fs} -> hub.Session -> hub.Agent
	r.HandleFunc("/api/grpc", h.handleGRPC).Methods("GET").
		Queries("id", "{id}")

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

	wsconn, err := wetty.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer wsconn.Close()
	rwc := wrap.WsConnToReadWriteCloser(wsconn)
	for range time.Tick(time.Second) {
		_, err := rwc.Write([]byte(pretty.JSONString(h.GetAgents())))
		if err != nil {
			log.Println(err)
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

	delegate := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		switch {
		case strings.HasPrefix(subpath, "/ws"):
			wsRelay(ag)(w, r)
		case strings.HasPrefix(subpath, "/rootfs"):
			fsRelay(ag)(w, r)
		case h.localui:
			staticFileHandler.ServeHTTP(w, r)
		default:
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
	})

	ag.BasicAuth(delegate).ServeHTTP(w, r)
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

		// common error: ws transport is closing
		log.Println(pipe(wrap.WsConnToReadWriteCloser(wsconn), sessionSendClient))
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

func (h *hub) handleGRPC(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		id   = vars["id"]
	)

	conn, err := wrap.Hijack(w)
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

// lys is an http handler to net.Listener converter
// lys implements net.Listener/http.Handler
type lys struct {
	conns chan net.Conn
}

func (l *lys) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := wrap.Hijack(w)
	if err != nil {
		return
	}
	l.conns <- conn
}

func (l *lys) Accept() (net.Conn, error) {
	return <-l.conns, nil
}

func (l *lys) Close() error {
	return nil
}

func (l *lys) Addr() net.Addr {
	return l
}

func (l *lys) Network() string {
	return "hijack"
}

func (l *lys) String() string {
	return l.Network()
}

func NewHandleHijackListener() *lys {
	return &lys{conns: make(chan net.Conn)}
}

func (h *hub) startYRPCServer() {
	var (
		listhand              = NewHandleHijackListener()
		handler  http.Handler = listhand
		listener net.Listener = listhand
	)
	log.Println(handler)
	log.Println(listener)
	h.handleYRPC = handler
	go h.serveYRPC(listener)
}

func ToRPC(conn net.Conn) types.RPC {
	return &YS{
		id:      uuid.New(),
		name:    rng.New(),
		created: time.Now(),
		Conn:    conn,
	}
}

func (ys *YS) Time() time.Time {
	return ys.created
}

func (ys *YS) Name() string {
	return ys.name
}

func (ys *YS) ID() string {
	return ys.id
}

func (ys *YS) RemoteIP() string {
	ip, _, _ := net.SplitHostPort(ys.Conn.RemoteAddr().String())
	return ip
}

type YS struct {
	id      string
	name    string
	created time.Time
	net.Conn
}

func (ys *YS) NewSession() {
	io.WriteString(ys.Conn, fmt.Sprintln("ACCEPT"))
}
