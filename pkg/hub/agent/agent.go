package agent

import (
	"net"
	"net/http"
	"time"

	auth "github.com/abbot/go-http-auth"
	"github.com/btwiuse/pretty"
	"k0s.io/pkg/api"
	"k0s.io/pkg/hub"
	"k0s.io/pkg/hub/agent/info"
)

var (
	_ hub.Agent = (*agent)(nil)
)

func NewAgent(session hub.Session, info *info.Info) hub.Agent {
	ag := &agent{
		info:             info,
		session:          session,
		created:          time.Now(),
		protocolHandlers: map[api.ProtocolID]chan net.Conn{},
	}

	ag.info.SetIP(session.RemoteIP())

	if *info.Auth {
		ag.htpasswd = info.Htpasswd
	}

	return ag
}

func (ag *agent) Info() *info.Info {
	return ag.info
}

func (ag *agent) MarshalJSON() ([]byte, error) {
	return []byte(pretty.JSONStringLine(ag.Info)), nil
}

type agent struct {
	info *info.Info // `json:"-"`

	protocolHandlers map[api.ProtocolID]chan net.Conn `json:"-"`
	session          hub.Session

	created  time.Time
	htpasswd map[string]string
}

func (ag *agent) OpenChannel(p api.ProtocolID) net.Conn {
	ag.session.OpenChannel(p)
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
	return ag.info.ID
}

func (ag *agent) Name() string {
	return ag.info.Name
}

func (ag *agent) ChannelChan(p api.ProtocolID) chan net.Conn {
	// ensure the channel is not nil
	_, ok := ag.protocolHandlers[p]
	if !ok {
		ag.protocolHandlers[p] = make(chan net.Conn)
	}
	return ag.protocolHandlers[p]
}
