package hub

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strings"
	"time"

	"github.com/btwiuse/invctrl/pkg/api"
	rpcimpl "github.com/btwiuse/invctrl/pkg/api/rpc/impl"
	"github.com/btwiuse/invctrl/wrap"
	"github.com/btwiuse/wetty/assets"
	"github.com/btwiuse/wetty/wetty"
	"github.com/gorilla/websocket"
	"github.com/kr/pty"
	"google.golang.org/grpc"
	"modernc.org/httpfs"
)

func getAgents(w http.ResponseWriter, r *http.Request) {
	GlobalAgentPool.Dump()
}

func grpcfactory(agent *Agent) (*grpc.ClientConn, error) {
	req := rpcimpl.NewSlaveRequest{
		Args: agent.Info,
	}
	resp := new(rpcimpl.NewSlaveResponse)

	done := make(chan *rpc.Call, 1)
	agent.RPCClient.Go("NewSlave.New", req, resp, done)
	<-done

	time.Sleep(time.Second / 3) // todo: investigate why it is necessary
	return <-agent.GRPCClientConn, nil
}

// listen on frontend ws connection
// serve http://<host>/id/ws endpoint
// see wslisten
// todo: translate between grpc/ws
func grpcrelayws(agent *Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", 405)
			return
		}

		// 1
		wsconn, err := wetty.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer wsconn.Close()

		// 2, 3
		cc, err := grpcfactory(agent)
		if err != nil {
			log.Println(err)
			return
		}

		// to bidiStreamClient
		bidiStreamClient := api.NewBidiStreamClient(cc)

		// to sendClient
		sendClient, err := bidiStreamClient.Send(context.Background())
		if err != nil {
			log.Fatalln(bidiStreamClient, err)
		}

		// go for send Input   <=	    websocket
		// go for send Resize  <= [wsrelay] websocket
		// for recv Output     =>	    websocket
		// (through chan Message{Type, Body} instead of interface)

		// todo: read from ws instead of generating fake msg
		go func() {
			for range time.Tick(time.Second) {
				rows, cols, err := pty.Getsize(os.Stdin)
				if err != nil {
					log.Println(err)
					break
				}

				json := fmt.Sprintf(`{"Cols": %d, "Rows": %d}`, cols, rows)

				resizeMsg := &api.Message{Type: []byte{wetty.ResizeTerminal}, Body: []byte(json)}
				err = sendClient.Send(resizeMsg)
				if err != nil {
					log.Println(err)
					return
				}
			}
		}()

		go func() {
			buf := make([]byte, 1024)
			for {
				n, err := os.Stdin.Read(buf)
				if err != nil {
					log.Println(err)
					break
				}

				inputMsg := &api.Message{Type: []byte{wetty.Input}, Body: buf[:n]}
				err = sendClient.Send(inputMsg)
				if err != nil {
					log.Println(err)
					continue
				}
			}
		}()

		for {
			select {
			// case <-ctx.Done():
			// log.Println("client dead:", c.RemoteAddr())
			// log.Println("client dead:")
			// return
			default:
				resp, err := sendClient.Recv()
				if err != nil {
					log.Println(err)
					return
				}
				fmt.Printf("%s", resp.Body)
			}
		}

	}
}

// here master and slave are connected via network
func Pipe(client, slave io.ReadWriteCloser) error {
	errs := make(chan error, 2)
	closeall := func() {
		client.Close()
		slave.Close()
	}

	go func() {
		log.Println("Pipe: client <= slave")
		defer closeall()
		_, err := io.Copy(client, slave)
		errs <- err
		// if you set it to 400, master will receive 399- bytes on each read
		// this value should at least CSPair.bufferSize + 1
		// otherwise some messages may be partially sent
	}()

	go func() {
		log.Println("Pipe: client => slave")
		defer closeall()
		_, err := io.Copy(slave, client)
		errs <- err
	}()

	return <-errs
}

// https://github.com/frankdejonge/golang-websocket-example/blob/master/connection.go
func handleConnectionError(slave io.Writer, err error) {
	ers := []int{
		// websocket.CloseNormalClosure,
		// websocket.CloseGoingAway,
		// websocket.CloseProtocolError,
		// websocket.CloseUnsupportedData,
		websocket.CloseNoStatusReceived,
		// websocket.CloseAbnormalClosure,
		// websocket.CloseInvalidFramePayloadData,
		// websocket.ClosePolicyViolation,
		// websocket.CloseMessageTooBig,
		// websocket.CloseMandatoryExtension,
		// websocket.CloseInternalServerErr,
		// websocket.CloseServiceRestart,
		// websocket.CloseTryAgainLater,
		// websocket.CloseTLSHandshake,
	}

	if websocket.IsUnexpectedCloseError(err, ers...) {
		log.Printf("Unexpected error from connection: %q", err)
	} else {
		log.Printf("Client close detected: %q, sending ClientDead message to slave", err)
	}

	if _, err := slave.Write([]byte{wetty.ClientDead}); err != nil {
		log.Println("Pipe:", err)
	}
}

/*
base :=
static:
$uuid
$uuid/css/*
$uuid/js/*
$uuid/png/*

grpcrelayws:
$uuid/ws

fsrelay:
$uuid/rootfs

new/slave:
/new/slave
*/
func static(w http.ResponseWriter, r *http.Request) {
	var staticFileServer, staticFileHandler http.Handler
	var id string
	base := strings.TrimPrefix(r.RequestURI, "/ws/")
	parts := strings.Split(base, "/")
	if len(parts) == 0 {
		goto HOME
	}
	id = parts[0]
	staticFileServer = http.FileServer(httpfs.NewFileSystem(assets.Assets, time.Now()))
	staticFileHandler = http.StripPrefix("/ws/"+id+"/", staticFileServer)

	if !GlobalAgentPool.Has(id) {
		goto HOME
	}

	if len(parts) == 1 {
		staticFileHandler.ServeHTTP(w, r)
		return
	}

	switch parts[1] {
	case "ws":
		// relayfrontend ws connection
		agent := GlobalAgentPool.Get(id)
		grpcrelayws(agent)(w, r)
	case "rootfs":
		// slave := GlobalAgentPool.Get(id)
		// p := "/" + path.Join(parts[2:]...)
		// fsrelay(w, r, slave, p)
	default:
		// handle non-ws static resources
		staticFileHandler.ServeHTTP(w, r)
	}
	return
HOME:
	http.Redirect(w, r, "/", 302)
}

func newAgent(w http.ResponseWriter, r *http.Request) {

	conn, err := wrap.WrapConn(w.(http.Hijacker).Hijack())
	if err != nil {
		log.Println("error hijacking:", err)
		return
	}

	agent := new(Agent) // using named return causes panic, why?
	agent.Info = r.URL.Query()
	agent.GRPCClientConn = make(chan *grpc.ClientConn)

	agent.RPCClient = agent.toRPCClient(conn)

	log.Println("new agent connected:", agent.Info.Get("id"))

	GlobalAgentPool.Add(agent)
}

func newSlave(w http.ResponseWriter, r *http.Request) {
	agent := GlobalAgentPool.Get(r.URL.Query().Get("id"))

	// extract conn
	conn, err := wrap.WrapConn(w.(http.Hijacker).Hijack())
	if err != nil {
		return // nil, err
	}

	// to client conn
	cc := toClientConn(conn)

	agent.GRPCClientConn <- cc
	return
}

func toClientConn(c net.Conn) *grpc.ClientConn {
	options := []grpc.DialOption{
		// grpc.WithTransportCredentials(creds),
		grpc.WithInsecure(),
		grpc.WithContextDialer(
			func(ctx context.Context, s string) (net.Conn, error) {
				return c, nil
			},
		),
	}

	cc, err := grpc.Dial("", options...)
	if err != nil {
		log.Fatal(err.Error())
	}
	return cc
}
