package agent

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/btwiuse/conntroll/pkg/api"
	"github.com/btwiuse/conntroll/pkg/hub"
	"github.com/btwiuse/conntroll/pkg/hub/agent/info"
	"github.com/btwiuse/conntroll/pkg/hub/session"
	"google.golang.org/grpc"
)

var (
	_ hub.Agent = (*agent)(nil)
)

func NewAgent(rpc hub.RPC, info hub.Info, xopts ...Opt) hub.Agent {
	ag := &agent{
		rpc:            rpc,
		SessionManager: NewSessionManager(),
		sch:            make(chan hub.Session),
		created:        time.Now(),
		Connected:      time.Now().Unix(),
		done:           make(chan struct{}),
		closeOnceDone:  &sync.Once{},
		IP:             rpc.RemoteIP(),
		Auth:		new(bool),
	}
	ag.fromInfo(info)

	for _, opt := range xopts {
		opt(ag)
	}

	return ag
}

func (ag *agent) fromInfo(info hub.Info) {
	ag.ID_ = info.GetID()
	ag.Name_ = info.GetName()
	ag.Tags = info.GetTags()
	if bahash := info.GetAuth(); bahash != "" {
		ag.bahash = bahash
		*ag.Auth = true
	}

	ag.OS = info.GetOS()
	ag.Pwd = info.GetPwd()
	ag.Arch = info.GetArch()
	ag.Distro = info.GetDistro()
	ag.Username = info.GetUsername()
	ag.Hostname = info.GetHostname()
}

type agent struct {
	hub.SessionManager `json:"-"`
	rpc                hub.RPC
	sch                chan hub.Session
	done               chan struct{}
	closeOnceDone      *sync.Once

	created time.Time
	bahash  string

	grpcCounter int
	rpcCounter  int

	info.Info `json:"-"` // inherit methods

	// Metadata, for flatten json output

	ID_       string   `json:"id"`
	Name_     string   `json:"name"`
	Tags      []string `json:"tags"`
	Auth      *bool    `json:"auth,omitempty"`
	Connected int64    `json:"connected"`
	IP        string   `json:"ip"`

	OS       string `json:"os"`
	Pwd      string `json:"pwd"`
	Arch     string `json:"arch"`
	Distro   string `json:"distro,omitempty"`
	Username string `json:"username"`
	Hostname string `json:"hostname"`
}

type Opt func(*agent)

func SetIP(ip string) Opt {
	return func(ag *agent) {
		ag.IP = ip
	}
}

func (ag *agent) newSession() {
	ag.rpc.NewSession()
}

func (ag *agent) NewSession() hub.Session {
	ag.newSession()
	return <-ag.sch
}

func (ag *agent) Close() {
	ag.closeOnceDone.Do(func() {
		close(ag.done)
	})
}

func (ag *agent) Done() <-chan struct{} {
	return ag.done
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

func (ag *agent) ID() string {
	return ag.ID_
}

func (ag *agent) Name() string {
	return ag.Name_
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
