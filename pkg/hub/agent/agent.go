package agent

import (
	"fmt"
	"net"
	"net/http"
	"time"

	auth "github.com/abbot/go-http-auth"
	"github.com/btwiuse/pretty"
	"k0s.io/k0s/pkg/hub"
	"k0s.io/k0s/pkg/hub/session"
)

var (
	_ hub.Agent = (*agent)(nil)
)

func NewAgent(rpc hub.RPC, info hub.AgentInfo) hub.Agent {
	ag := &agent{
		rpc:            rpc,
		created:        time.Now(),
		SessionManager: NewSessionManager(),
		// TerminalManager: NewTerminalManager(),
		termch:    make(chan net.Conn),
		sch:       make(chan hub.Session),
		socks5ch:  make(chan net.Conn),
		redirch:   make(chan net.Conn),
		fsch:      make(chan net.Conn),
		metricsch: make(chan net.Conn),
	}

	info.SetIP(rpc.RemoteIP())
	ag.AgentInfo = info

	if info.GetAuth() {
		ag.htpasswd = info.GetHtpasswd()
	}

	return ag
}

func (ag *agent) MarshalJSON() ([]byte, error) {
	return []byte(pretty.JSONString(ag.AgentInfo)), nil
}

type agent struct {
	hub.AgentInfo // `json:"-"` // inherit methods

	hub.SessionManager `json:"-"`
	// hub.TerminalManager `json:"-"`
	rpc       hub.RPC
	sch       chan hub.Session
	socks5ch  chan net.Conn
	redirch   chan net.Conn
	fsch      chan net.Conn
	termch    chan net.Conn
	metricsch chan net.Conn

	created  time.Time
	htpasswd map[string]string

	grpcCounter int
}

func (ag *agent) NewSocks5() net.Conn {
	ag.rpc.NewSocks5()
	return <-ag.socks5ch
}

func (ag *agent) NewRedir() net.Conn {
	ag.rpc.NewSocks5()
	return <-ag.redirch
}

func (ag *agent) NewFS() net.Conn {
	ag.rpc.NewFS()
	return <-ag.fsch
}

func (ag *agent) NewMetrics() net.Conn {
	ag.rpc.NewMetrics()
	return <-ag.metricsch
}

func (ag *agent) NewSession() hub.Session {
	ag.rpc.NewSession()
	return <-ag.sch
}

func (ag *agent) NewTerminal() net.Conn {
	ag.rpc.NewTerminal()
	return <-ag.termch
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
	return ag.GetID()
}

func (ag *agent) Name() string {
	return ag.GetName()
}

// blocks until agent.NewFS reads the channel
func (ag *agent) AddFSConn(conn net.Conn) {
	ag.fsch <- conn
}

// blocks until agent.NewSocks5 reads the channel
func (ag *agent) AddSocks5Conn(conn net.Conn) {
	ag.socks5ch <- conn
}

// blocks until agent.NewRedir reads the channel
func (ag *agent) AddRedirConn(conn net.Conn) {
	ag.redirch <- conn
}

// blocks until agent.NewMetrics reads the channel
func (ag *agent) AddMetricsConn(conn net.Conn) {
	ag.metricsch <- conn
}

// blocks until agent.NewSession reads the channel
func (ag *agent) AddSessionConn(conn net.Conn) {
	ag.grpcCounter += 1
	name := fmt.Sprintf("%s.%d", ag.Name(), ag.grpcCounter)
	// ag.sch <- session.NewSession(name, api.NewSessionClient(toGRPCClientConn(c)))
	ag.sch <- session.NewSession(name, conn)
}

// blocks until agent.NewTerminal reads the channel
func (ag *agent) AddTerminalConn(conn net.Conn) {
	ag.termch <- conn
}
