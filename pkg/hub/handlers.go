package hub

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/url"
	"strings"
	"time"

	"github.com/btwiuse/invctrl/pkg/api"
	rpcimpl "github.com/btwiuse/invctrl/pkg/api/rpc/impl"
	"github.com/btwiuse/invctrl/pkg/wrap"
	"github.com/btwiuse/pretty"
	"github.com/btwiuse/wetty/assets"
	"github.com/btwiuse/wetty/utils"
	"github.com/btwiuse/wetty/wetty"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"modernc.org/httpfs"
)

func getAgents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	infos := []url.Values{}
	for _, v := range GlobalAgentPool.Agents.Values() {
		slave := v.(*Agent)
		infos = append(infos, slave.Info)
	}
	w.Write([]byte(pretty.JSONString(infos)))
}

func grpcfactory(agent *Agent) (*grpc.ClientConn, error) {
	req := rpcimpl.NewSlaveRequest{
		Info: agent.Info,
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
		bidiStreamClient := api.NewSlaveClient(cc)

		// to sendClient
		slave, err := bidiStreamClient.Send(context.Background())
		if err != nil {
			log.Fatalln(slave, err)
		}

		pipe(utils.WsConnToReadWriter(wsconn), slave)
	}
}

// here master and slave are connected via network
// go for send Input   <=	   websocket
// go for send Resize  <= [ Pipe ] websocket
// for recv Output     =>	   websocket
// (through chan Message{Type, Body} instead of interface)
func pipe(ws io.ReadWriteCloser, slave api.Slave_SendClient) error {
	errs := make(chan error, 2)
	closeall := func() {
		ws.Write([]byte{wetty.SlaveDead})
		slave.Send(&api.Message{Type: []byte{wetty.ClientDead}})
		ws.Close()
		slave.CloseSend()
	}

	go func() {
		log.Println("pipe: client(ws) => slave(grpc)")
		defer closeall()
		buf := make([]byte, 1<<12)
		for {
			n, err := ws.Read(buf)
			if err != nil {
				errs <- err
				break
			}
			msg := &api.Message{Type: buf[:1], Body: buf[1:n]}
			err = slave.Send(msg)
			if err != nil {
				errs <- err
				return
			}
		}
	}()

	go func() {
		log.Println("pipe: client(ws) <= slave(grpc)")
		defer closeall()
		for {
			resp, err := slave.Recv()
			if err != nil {
				errs <- err
				break
			}
			_, err = ws.Write(append(resp.Type, resp.Body...))
			if err != nil {
				errs <- err
				break
			}
		}
	}()

	firstError := <-errs
	log.Println(firstError)
	return firstError
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
		log.Println("pipe:", err)
	}
}

func static(w http.ResponseWriter, r *http.Request) {
	var staticFileServer, staticFileHandler http.Handler
	var id string
	base := strings.TrimPrefix(r.RequestURI, "/agent/")
	parts := strings.Split(base, "/")
	if len(parts) == 0 {
		goto HOME
	}
	id = parts[0]
	staticFileServer = http.FileServer(httpfs.NewFileSystem(assets.Assets, time.Now()))
	staticFileHandler = http.StripPrefix("/agent/"+id+"/", staticFileServer)

	if !GlobalAgentPool.Has(id) {
		goto HOME
	}

	if len(parts) == 1 {
		staticFileHandler.ServeHTTP(w, r)
		return
	}

	switch parts[1] {
	case "ws":
		grpcrelayws(GlobalAgentPool.Get(id))(w, r)
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

func newAgentSlave(w http.ResponseWriter, r *http.Request) {
	conn, err := wrap.WrapConn(w.(http.Hijacker).Hijack())
	if err != nil {
		log.Println("error hijacking:", err)
		return
	}
	id := r.URL.Query().Get("id")

	if !GlobalAgentPool.Has(id) {
		agent := &Agent{
			Info:           r.URL.Query(),
			GRPCClientConn: make(chan *grpc.ClientConn),
		}
		agent.MakeInterceptedRPCClient(conn)

		GlobalAgentPool.Add(agent)
		log.Println("new agent connected:", agent.Info.Get("id"))
		return
	}

	agent := GlobalAgentPool.Get(id)
	agent.GRPCClientConn <- toGRPCClientConn(conn)
	log.Println("new slave connected:", agent.Info.Get("id"))
}

func toGRPCClientConn(c net.Conn) *grpc.ClientConn {
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
