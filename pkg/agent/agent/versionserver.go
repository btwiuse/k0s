package agent

import (
	"net"
	"net/http"

	"github.com/btwiuse/pretty"
	types "k0s.io/k0s/pkg/agent"
)

func StartVersionServer(c types.Config) chan net.Conn {
	var (
		versionListener = NewLys()
		handler         = LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(pretty.JSONString(c.GetVersion())))
		}))
		versionServer = &http.Server{Handler: handler}
	)
	go versionServer.Serve(versionListener)
	return versionListener.Conns
}
