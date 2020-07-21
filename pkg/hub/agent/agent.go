package agent

import (
	"net"
	"net/http"
	"time"

	auth "github.com/abbot/go-http-auth"
	"github.com/btwiuse/pretty"
	"k0s.io/k0s/pkg/api"
	"k0s.io/k0s/pkg/hub"
)

var (
	_ hub.Agent = (*agent)(nil)
)

func newTunnels() (tuns map[api.Tunnel]chan net.Conn) {
	for i := 0; i < 256; i++ {
		tuns[api.Tunnel(i)] = make(chan net.Conn)
	}
	return
}

func NewAgent(rpc hub.RPC, info hub.AgentInfo) hub.Agent {
	ag := &agent{
		rpc:       rpc,
		created:   time.Now(),
		Tunnels:   newTunnels(),
		grpcNames: make(map[net.Conn]string),
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

	Tunnels map[api.Tunnel]chan net.Conn `json:"-"`
	rpc     hub.RPC

	created  time.Time
	htpasswd map[string]string

	grpcCounter int
	grpcNames   map[net.Conn]string
}

func (ag *agent) NewTunnel(tun api.Tunnel) net.Conn {
	ag.rpc.NewTunnel(tun)
	return <-ag.Tunnels[tun]
}

func (ag *agent) BasicAuth(next http.Handler) http.Handler {
	if len(ag.htpasswd) == 0 {
		return next
	}
	secret := func(user, realm string) string {
		// realm = "please enter password for " + ag.Name()
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

// blocks until agent.NewTunnel(api.Tunnel) reads the channel
func (ag *agent) AddTunnel(tun api.Tunnel, conn net.Conn) {
	ag.Tunnels[tun] <- conn
}
