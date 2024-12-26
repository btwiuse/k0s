package agent

import (
	"net"
	"net/http"
	"time"

	auth "github.com/abbot/go-http-auth"
	"github.com/btwiuse/pretty"
	"k0s.io/pkg/api"
	"k0s.io/pkg/hub"
)

var (
	_ hub.Agent = (*agent)(nil)
)

func NewAgent(rpc hub.RPC, info hub.AgentInfo) hub.Agent {
	ag := &agent{
		rpc:       rpc,
		created:   time.Now(),
		channels:  map[api.ProtocolID]chan net.Conn{},
	}

	info.SetIP(rpc.RemoteIP())
	ag.AgentInfo = info

	if info.GetAuth() {
		ag.htpasswd = info.GetHtpasswd()
	}

	return ag
}

func (ag *agent) MarshalJSON() ([]byte, error) {
	return []byte(pretty.JSONStringLine(ag.AgentInfo)), nil
}

type agent struct {
	hub.AgentInfo // `json:"-"` // inherit methods

	channels map[api.ProtocolID]chan net.Conn `json:"-"`
	rpc      hub.RPC

	created  time.Time
	htpasswd map[string]string
}

func (ag *agent) OpenChannel(p api.ProtocolID) net.Conn {
	ag.rpc.OpenChannel(p)
	return <-ag.ChannelChan(p)
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

func (ag *agent) ChannelChan(p api.ProtocolID) chan net.Conn {
	// ensure the channel is not nil
	_, ok := ag.channels[p]
	if !ok {
		ag.channels[p] = make(chan net.Conn)
	}
	return ag.channels[p]
}
