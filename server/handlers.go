package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/rpc"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/navigaid/gotty/assets"
	"github.com/navigaid/gotty/utils"
	"github.com/navigaid/gotty/wetty"
	"modernc.org/httpfs"

	"github.com/invctrl/hijack/protocol"
	"github.com/invctrl/hijack/wrap"
)

func hijack(original http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(http.CanonicalHeaderKey("Hijack")) == "true" {
			hijacker(w, r)
			return
		}
		original.ServeHTTP(w, r)
	}
}

func hijacker(w http.ResponseWriter, r *http.Request) {
	_, err := NewClient(w)
	if err != nil {
		log.Println(err)
		return
	}
}

// rpc call to initiate a ws conn to /ws from client
func wsfactory(client *Client) (*websocket.Conn, error) {
	nonce := uuid.New().String()
	id := client.UUID.String()

	req := protocol.WsConnRequest{
		Id:    id,
		Nonce: nonce,
	}
	resp := new(protocol.WsConnResponse)

	done := make(chan *rpc.Call, 1)
	client.RPC.Go("WsConn.New", req, resp, done)
	<-done

	time.Sleep(time.Second / 3) // todo: investigate why it is necessary
	conn, ok := client.WsConns[nonce]
	if !ok {
		return nil, fmt.Errorf("client nonce doesn't exist: %s", nonce)
	}
	delete(client.WsConns, nonce)
	return conn, nil
}

// wslisten assumes r.RequestURI == "/ws" {
// the request is initiated by slaves at wsfactory's command
// call chain:
// 1. browser (ws:conn)-> ws://<host>/id/ws
// 2. host (rpc on tcp)-> rpc://slave.WsConn.New(id, nounce)
// 3. slave (ws:connRemote)-> ws://<host>/ws (with id and nounce to identify slave)
// 4. browser <=(conn:ws:connRemote)=> slave
func wslisten(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	conn, err := wetty.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// don't do this!!!
	// defer conn.Close()
	// you should leave the conn open until the matching frontend conn is closed

	var this io.ReadWriter = &utils.WsWrapper{conn}

	buf := make([]byte, 1024)
	n, err := this.Read(buf)
	if err != nil {
		log.Println(err)
	}
	id := string(buf[:n])
	client := ClientPool.Get(id)

	n, err = this.Read(buf)
	if err != nil {
		log.Println(err)
	}
	nonce := string(buf[:n])

	client.WsConns[nonce] = conn
}

// relay fs http request
// drawback: poor performance handling large files
// 0. rewrite request path
// 1. record client request (httputil.DumpRequest)
// (2. send recorded request to grpc
// (3. record and dump grpc response (httptest.ResponseRecorder, httputil.DumpResponse)
// (4. return grpc response
// 5(theoretically). decode grpc response (http.ReadResponse)
// 5(practically). hijack w.(http.Hijacker).Hijack()...Write(grpc.Response)
func fsrelay(w http.ResponseWriter, r *http.Request, client *Client, p string) {
	log.Println("request uri:", r.RequestURI)

	// 0.
	r.RequestURI = strings.SplitN(r.RequestURI, "rootfs", 2)[1]

	log.Println("path:", r.RequestURI)

	// 1.
	reqbuf, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println(err)
		return
	}

	// 2.
	req := protocol.RootfsRequest{
		Request: reqbuf,
	}
	resp := new(protocol.RootfsResponse)

	log.Println("calling Rootfs.New")

	if err := client.RPC.Call("Rootfs.New", req, resp); err != nil {
		log.Println(err)
		return
	}

	// 5. hijack client http request -> net.Conn
	rwc, err := wrap.WrapConn(w.(http.Hijacker).Hijack())
	if err != nil {
		log.Println(err)
		return
	}

	defer rwc.Close() // tell browser to stop loading

	if _, err := rwc.Write(resp.Response); err != nil {
		log.Println(err)
		return
	}
}

// listen on frontend ws connection
// serve http://<host>/id/ws endpoint
// see wslisten
func wsrelay(w http.ResponseWriter, r *http.Request, client *Client) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	// 1
	conn, err := wetty.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// 2, 3
	connRemote, err := wsfactory(client)
	if err != nil {
		log.Println(err)
		return
	}

	var (
		this io.ReadWriter = &utils.WsWrapper{conn}
		that io.ReadWriter = &utils.WsWrapper{connRemote}
		errs               = make(chan error, 2)
	)

	// 4
	go func() {
		errs <- func() error {
			_, err := io.Copy(that, this)
			log.Println(err)
			return err
		}()
	}()

	go func() {
		errs <- func() error {
			_, err := io.Copy(this, that)
			log.Println(err)
			return err
		}()
	}()

	log.Println(<-errs)

	return
}

/*
base :=
$uuid
$uuid/css/*
$uuid/js/*
$uuid/png/*

$uuid/ws

$uuid/rootfs
*/
func frontend(w http.ResponseWriter, r *http.Request) {
	var staticFileServer, staticFileHandler http.Handler
	var id string
	base := strings.TrimPrefix(r.RequestURI, "/ws/")
	parts := strings.Split(base, "/")
	if len(parts) == 0 {
		goto REDIR
	}
	id = parts[0]
	staticFileServer = http.FileServer(httpfs.NewFileSystem(assets.Assets, time.Now()))
	staticFileHandler = http.StripPrefix("/ws/"+id+"/", staticFileServer)

	if !ClientPool.Has(id) {
		goto REDIR
	}

	if len(parts) == 1 {
		staticFileHandler.ServeHTTP(w, r)
		return
	}

	switch parts[1] {
	case "ws":
		// relayfrontend ws connection
		client := ClientPool.Get(id)
		wsrelay(w, r, client)
	case "rootfs":
		client := ClientPool.Get(id)
		p := "/" + path.Join(parts[2:]...)
		fsrelay(w, r, client, p)
	default:
		// handle non-ws static resources
		staticFileHandler.ServeHTTP(w, r)
	}
	return
REDIR:
	http.Redirect(w, r, "/", 302)
}

func ls(w http.ResponseWriter, r *http.Request) {
	isCurrent := func(uuid string) string {
		if (ClientPool.Current != nil) && (ClientPool.Current.UUID.String() == uuid) {
			return "*"
		}
		return " "
	}

	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	uri := strings.TrimPrefix(r.RequestURI, "/")
	if ClientPool.Has(uri) {
		log.Println(uri)
		client := ClientPool.Get(uri)
		href := fmt.Sprintf(`<a href="/ws/%s/">%s</a>`, uri, uri)
		fmt.Fprintln(w, href)
		fmt.Fprintln(w, "<pre>")
		fmt.Fprintln(w,
			fmt.Sprintf("[%s]", "N/A"),
			isCurrent(uri),
			uri,
			"ssh ubuntu@"+strings.Split(client.RemoteAddr.String(), ":")[0],
			client.Info,
		)
		fmt.Fprintln(w, "</pre>")
		return
	}
	for i, v := range ClientPool.Clients.Values() {
		client := v.(*Client)
		uuid := ClientPool.Clients.Keys()[i].(string)
		href := fmt.Sprintf(`<a href="%s">%s</a>`, uuid, uuid)
		wshref := fmt.Sprintf(`<a href="/ws/%s/">ws</a>`, uuid)
		fshref := fmt.Sprintf(`<a href="/ws/%s/rootfs/">fs</a>`, uuid)
		fmt.Fprintln(w,
			fmt.Sprintf("[%s]", strconv.Itoa(i+1)),
			isCurrent(uuid),
			href,
			wshref,
			fshref,
			"ssh ubuntu@"+strings.Split(client.RemoteAddr.String(), ":")[0],
		)
		fmt.Fprintln(w, "<pre>")
		fmt.Fprintln(w,
			client.Info,
		)
		fmt.Fprintln(w, "</pre>")
	}
}
