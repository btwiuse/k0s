package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"
	"time"

	"modernc.org/httpfs"
	//"github.com/davecgh/go-spew/spew"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/invctrl/hijack/protocol"
	"github.com/navigaid/gotty/assets"
	"github.com/navigaid/gotty/utils"
	"github.com/navigaid/gotty/wetty"
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

// listen on frontend ws connection
func wsrelay(w http.ResponseWriter, r *http.Request, client *Client) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	conn, err := wetty.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

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

func frontend(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(strings.TrimPrefix(r.RequestURI, "/ws/"), "/")[0]

	if ClientPool.Has(uri) {
		client := ClientPool.Get(uri)
		// handle non-ws static resources
		if !strings.HasSuffix(r.RequestURI, "/ws") {
			staticFileServer := http.FileServer(httpfs.NewFileSystem(assets.Assets, time.Now()))
			http.StripPrefix("/ws/"+uri+"/", staticFileServer).ServeHTTP(w, r)
			return
		}
		// relayfrontend ws connection
		wsrelay(w, r, client)
		return
	}
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
		fmt.Fprintln(w,
			fmt.Sprintf("[%s]", strconv.Itoa(i+1)),
			isCurrent(uuid),
			href,
			wshref,
			"ssh ubuntu@"+strings.Split(client.RemoteAddr.String(), ":")[0],
		)
		fmt.Fprintln(w, "<pre>")
		fmt.Fprintln(w,
			client.Info,
		)
		fmt.Fprintln(w, "</pre>")
	}
}
