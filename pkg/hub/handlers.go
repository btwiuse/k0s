package hub

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/btwiuse/conntroll/pkg/api"
	"github.com/btwiuse/conntroll/pkg/wrap"
	"github.com/btwiuse/pretty"
	"github.com/btwiuse/wetty/pkg/assets"
	"github.com/btwiuse/wetty/pkg/msg"
	"github.com/btwiuse/wetty/pkg/utils"
	"github.com/btwiuse/wetty/pkg/wetty"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"modernc.org/httpfs"
)

func getAgents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	agents := []*Agent{}
	for _, v := range GlobalAgentPool.Values() {
		agents = append(agents, v.(*Agent))
	}
	w.Write([]byte(pretty.JSONString(agents)))
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

		session := NewSession(agent)
		sessionSendClient, err := session.Send(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		pipe(utils.WsConnToReadWriter(wsconn), sessionSendClient)
	}
}

// here client and session are connected via hub
// go for send Input   <=	   websocket
// go for send Resize  <= [ Pipe ] websocket
// for recv Output     =>	   websocket
// (through chan Message{Type, Body} instead of interface)
func pipe(ws io.ReadWriteCloser, session api.Session_SendClient) error {
	defer ws.Close()
	errs := make(chan error, 2)

	go func() {
		log.Println("pipe: client(ws) => session(grpc)")
		buf := make([]byte, 1<<12)
		for {
			n, err := ws.Read(buf)
			if err != nil {
				errs <- err // client closed
				break
			}
			msg := &api.Message{Type: msg.Type(buf[0]), Body: buf[1:n]}
			err = session.Send(msg)
			if err != nil {
				errs <- err // session closed
				return
			}
		}
	}()

	go func() {
		log.Println("pipe: client(ws) <= session(grpc)")
		for {
			resp, err := session.Recv()
			if err != nil {
				errs <- err // session closed
				break
			}
			_, err = ws.Write(append([]byte{byte(resp.Type)}, resp.Body...))
			if err != nil {
				errs <- err // client closed
				break
			}
		}
	}()

	firstError := <-errs
	log.Println(firstError)
	return firstError
}

// https://github.com/frankdejonge/golang-websocket-example/blob/master/connection.go
func handleConnectionError(session io.Writer, err error) {
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
		log.Printf("Client close detected: %q, sending ClientDead message to session", err)
	}

	// if _, err := session.Write([]byte{msg.Type_SESSION_CLOSE}); err != nil {
	//	log.Println("pipe:", err)
	// }
}

func static(w http.ResponseWriter, r *http.Request) {
	var staticFileServer, staticFileHandler http.Handler
	var id string
	base := strings.TrimPrefix(r.RequestURI, "/agent/")
	parts := strings.Split(base, "/")
	if len(parts) == 0 {
		http.Redirect(w, r, "/", 302)
		return
	}
	id = parts[0]
	staticFileServer = http.FileServer(httpfs.NewFileSystem(assets.Assets, time.Now()))
	staticFileHandler = http.StripPrefix("/agent/"+id+"/", staticFileServer)

	if !GlobalAgentPool.Has(id) {
		http.Redirect(w, r, "/", 302)
		return
	}

	if len(parts) == 1 {
		staticFileHandler.ServeHTTP(w, r)
		return
	}

	agent := GlobalAgentPool.Get(id)

	switch parts[1] {
	case "ws":
		grpcrelayws(agent)(w, r)
	case "rootfs":
		fsrelay(agent)(w, r)
	default:
		// handle non-ws static resources
		staticFileHandler.ServeHTTP(w, r)
	}
}

func fsrelay(agent *Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 0.
		path := strings.SplitN(r.RequestURI, "rootfs", 2)[1]
		r.RequestURI = path

		// 1.
		reqbuf, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
			return
		}

		chunkRequest := &api.ChunkRequest{
			Path:    path,
			Request: reqbuf,
		}
		session := NewSession(agent)
		sessionChunkerClient, err := session.Chunker(context.Background(), chunkRequest)
		if err != nil {
			log.Println(err)
			return
		}

		// 5. hijack slave http request -> net.Conn
		rwc, err := wrap.WrapConn(w.(http.Hijacker).Hijack())
		if err != nil {
			log.Println(err)
			return
		}

		defer rwc.Close() // tell browser to stop loading

		for {
			chunk, err := sessionChunkerClient.Recv()
			if err != nil {
				if err != io.EOF {
					log.Println(err)
				}
				break
			}
			if _, err := rwc.Write(chunk.Chunk); err != nil {
				log.Println(err)
				break
			}
		}

		return
	}
}

func newAgentOrSession(w http.ResponseWriter, r *http.Request) {
	conn, err := wrap.WrapConn(w.(http.Hijacker).Hijack())
	if err != nil {
		log.Println("error hijacking:", err)
		return
	}
	id := r.URL.Query().Get("id")

	// new agent
	if !GlobalAgentPool.Has(id) {
		// TODO: validate request, ensure parameters are sane
		values := r.URL.Query()
		ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
		agent := &Agent{
			GRPCClientConn: make(chan *grpc.ClientConn),
			// Meta
			Id:        values["id"][0],
			IP:        ip,
			Pwd:       values["pwd"][0],
			Whoami:    values["whoami"][0],
			Hostname:  values["hostname"][0],
			Connected: time.Now().Unix(),
		}
		if len(values["os"]) != 0 {
			agent.OS = values["os"][0]
		}
		if len(values["arch"]) != 0 {
			agent.ARCH = values["arch"][0]
		}
		agent.MakeInterceptedRPCClient(conn)

		GlobalAgentPool.Add(agent)
		log.Println("new agent connected:", agent.Id)
		return
	}

	// new session
	agent := GlobalAgentPool.Get(id)
	agent.GRPCClientConn <- toGRPCClientConn(conn)
	log.Println("new session created for agent", agent.Id)
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
