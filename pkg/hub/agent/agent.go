package agent

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"
	"time"

	"github.com/btwiuse/conntroll/pkg/api"
	rpcimpl "github.com/btwiuse/conntroll/pkg/api/rpc/impl"
	"github.com/btwiuse/conntroll/pkg/hub"
	"github.com/btwiuse/conntroll/pkg/hub/session"
	"github.com/btwiuse/conntroll/pkg/uuid"
	"github.com/btwiuse/conntroll/pkg/wrap"
	"google.golang.org/grpc"
)

var (
	_ hub.Agent = (*agent)(nil)
)

func NewAgent(conn net.Conn, opts ...Opt) hub.Agent {
	ag := &agent{
		sch:            make(chan hub.Session, 5),
		SessionManager: NewSessionManager(),
		rpcManager:     NewRPCManager(),
		created:        time.Now(),
		Connected:      time.Now().Unix(),
		done:           make(chan struct{}),
		once:           &sync.Once{},
	}
	for _, opt := range opts {
		opt(ag)
	}
	ag.AddRPCConn(conn)
	for i := 0; i < 3; i++ {
		ag.NewRPC()
	}
	for i := 0; i < 5; i++ {
		ag.newSession()
	}
	return ag
}

type Opt func(*agent)

func SetTags(tags []string) Opt {
	return func(ag *agent) {
		ag.Tag = tags
	}
}

func SetDistro(dist string) Opt {
	return func(ag *agent) {
		ag.Distro = dist
	}
}

func SetName(name string) Opt {
	return func(ag *agent) {
		ag.Nam = name
	}
}

func SetID(id string) Opt {
	return func(ag *agent) {
		ag.Id = id
	}
}

func SetUsername(u string) Opt {
	return func(ag *agent) {
		ag.Usernam = u
	}
}

func SetBasicAuthHash(bahash string) Opt {
	return func(ag *agent) {
		ag.bahash = bahash
		ag.Auth = true
	}
}

func SetIP(ip string) Opt {
	return func(ag *agent) {
		ag.IP = ip
	}
}

func SetPWD(p string) Opt {
	return func(ag *agent) {
		ag.PWD = p
	}
}

func SetHostname(h string) Opt {
	return func(ag *agent) {
		ag.Hostnam = h
	}
}

func SetOS(o string) Opt {
	return func(ag *agent) {
		ag.OS = o
	}
}

func SetARCH(a string) Opt {
	return func(ag *agent) {
		ag.ARCH = a
	}
}

// make server call AddRPCConn
// this func is asynchronuous, we don't care the result
func (ag *agent) NewRPC() {
	req := rpcimpl.RPCRequest{}
	resp := &rpcimpl.RPCResponse{}
	rc := ag.rpcManager.Last()
	if rc == nil {
		log.Println("client dead:", ag.ID())
		ag.Close()
		return
	}
	err := rc.Call("RPC.New", req, resp)
	if err != nil {
		log.Println("RPC.New", err)
	}
}

func (ag *agent) newSession() {
	req := rpcimpl.SessionRequest{}
	resp := &rpcimpl.SessionResponse{}
	rc := ag.rpcManager.Last()
	if rc == nil {
		log.Println("client dead:", ag.ID())
		ag.Close()
		return
	}
	rc.Go("Session.New", req, resp, nil)
}

func (ag *agent) NewSession() hub.Session {
	if len(ag.sch) < 3 {
		ag.newSession()
		ag.newSession()
	}
	return <-ag.sch
}

func (ag *agent) Close() {
	ag.once.Do(func() {
		close(ag.done)
	})
}

func (ag *agent) Done() <-chan struct{} {
	return ag.done
}

type agent struct {
	hub.SessionManager `json:"-"`
	rpcManager         hub.RPCManager
	sch                chan hub.Session
	done               chan struct{}
	once               *sync.Once

	created time.Time
	bahash  string

	grpcCounter int
	rpcCounter  int

	// Metadata
	Id        string   `json:"id"`
	Nam       string   `json:"name"`
	Tag       []string `json:"tags"` // ,omitempty
	Connected int64    `json:"connected"`
	Hostnam   string   `json:"hostname"`
	Usernam   string   `json:"username"`
	PWD       string   `json:"pwd"`
	OS        string   `json:"os"`
	Distro    string   `json:"distro,omitempty"`
	ARCH      string   `json:"arch"`
	IP        string   `json:"ip"`
	Auth      bool     `json:"auth"`
}

func (ag *agent) BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ag.bahash != "" {
			username, password, ok := r.BasicAuth()
			uphash := fmt.Sprintf("%x", sha256.Sum256([]byte(username+":"+password)))
			log.Println(uphash, ag.bahash)
			if (!ok) || (uphash != ag.bahash) {
				realm := "please enter password for " + ag.Name()
				w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
				w.WriteHeader(401)
				w.Write([]byte("Unauthorised.\n"))
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (ag *agent) Time() time.Time {
	return ag.created
}

func (ag *agent) Name() string {
	return ag.Nam
}

func (ag *agent) ID() string {
	return ag.Id
}

func (ag *agent) Tags() []string {
	return ag.Tag
}

func (ag *agent) Username() string {
	return ag.Usernam
}

func (ag *agent) Hostname() string {
	return ag.Hostnam
}

// we use NewRPCClient over rpc.NewClient(conn)
// so we can remove agent from pool immediately when it is disconnected

// when we have multiple RPC clients, things man change a bit
// we may assume the agent is always connected now
// we should probably spawn another RPCClient onClose, instead of simply letting it die
func (ag *agent) AddRPCConn(c net.Conn) {
	id := uuid.New()
	pr, pw := io.Pipe()
	go func() {
		defer ag.onRPCClose(id)
		if _, err := io.Copy(pw, c); err != nil {
			log.Println(err)
		}
	}()
	rpcClient := rpc.NewClient(&wrap.ReadWriteCloser{
		Reader: pr,
		Writer: c,
		Closer: c,
	})
	ag.rpcCounter += 1
	name := fmt.Sprintf("%s-%d", ag.Name(), ag.rpcCounter)
	rc := ToRPC(name, id)(rpcClient)
	ag.rpcManager.AddRPC(rc)
}

// onclose is called when rpc connection is lost
func (ag *agent) onRPCClose(id string) {
	log.Println("disconnected:", ag.ID(), id)
	ag.rpcManager.Del(id)
	// assuming there are other rpc conns left
	ag.NewRPC()
}

// blocks until agent.NewSession reads the channel
func (ag *agent) AddSessionConn(c net.Conn) {
	toGRPCClientConn := func(c net.Conn) *grpc.ClientConn {
		options := []grpc.DialOption{
			// grpc.WithTransportCredentials(creds),
			grpc.WithInsecure(),
			grpc.WithContextDialer(
				func(ctx context.Context, s string) (net.Conn, error) {
					return c, nil
				},
			),
		}

		// TODO: handle this
		cc, err := grpc.Dial("", options...)
		if err != nil {
			log.Fatal(err.Error())
		}
		return cc
	}
	ag.grpcCounter += 1
	name := fmt.Sprintf("%s.%d", ag.Name(), ag.grpcCounter)
	ag.sch <- session.NewSession(name, api.NewSessionClient(toGRPCClientConn(c)))
}
