package agent

import (
	"fmt"
	"net"
	"net/http"
	"time"

	auth "github.com/abbot/go-http-auth"
	"k0s.io/conntroll/pkg/hub"
	"k0s.io/conntroll/pkg/hub/agent/info"
	"k0s.io/conntroll/pkg/hub/session"
)

var (
	_ hub.Agent = (*agent)(nil)
)

func NewAgent(rpc hub.RPC, info hub.Info, xopts ...Opt) hub.Agent {
	ag := &agent{
		rpc:            rpc,
		SessionManager: NewSessionManager(),
		sch:            make(chan hub.Session),
		socks5ch:       make(chan net.Conn),
		fsch:           make(chan net.Conn),
		created:        time.Now(),
		Connected:      time.Now().Unix(),
		IP:             rpc.RemoteIP(),
		Auth:           new(bool),
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

	if htpasswd := info.GetHtpasswd(); len(htpasswd) != 0 {
		ag.htpasswd = htpasswd
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
	socks5ch           chan net.Conn
	fsch               chan net.Conn

	created  time.Time
	htpasswd map[string]string

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

func (ag *agent) NewSocks5() net.Conn {
	ag.rpc.NewSocks5()
	return <-ag.socks5ch
}

func (ag *agent) NewFS() net.Conn {
	ag.rpc.NewFS()
	return <-ag.fsch
}

func (ag *agent) NewSession() hub.Session {
	ag.rpc.NewSession()
	return <-ag.sch
}

func (ag *agent) BasicAuth(next http.Handler) http.Handler {
	if len(ag.htpasswd) == 0 {
		return next
	}
	secret := func(user, realm string) string {
		realm = "please enter password for " + ag.Name()
		for k, v := range ag.htpasswd {
			if user == k {
				return v
			}
		}
		return ""
	}
	authenticator := auth.NewBasicAuthenticator("", secret)
	nextHandlerFunc := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
	return auth.JustCheck(authenticator, nextHandlerFunc)
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

// blocks until agent.NewFS reads the channel
func (ag *agent) AddFSConn(conn net.Conn) {
	ag.fsch <- conn
}

// blocks until agent.NewSocks5 reads the channel
func (ag *agent) AddSocks5Conn(conn net.Conn) {
	ag.socks5ch <- conn
}

// blocks until agent.NewSession reads the channel
func (ag *agent) AddSessionConn(conn net.Conn) {
	ag.grpcCounter += 1
	name := fmt.Sprintf("%s.%d", ag.Name(), ag.grpcCounter)
	// ag.sch <- session.NewSession(name, api.NewSessionClient(toGRPCClientConn(c)))
	ag.sch <- session.NewSession(name, conn)
}
