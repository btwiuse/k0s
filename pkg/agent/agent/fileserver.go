package agent

import (
	"net/http"

	types "k0s.io/conntroll/pkg/agent"
)

func StartFileServer(c types.Config) types.FileServer {
	fsListener := NewLys()
	fileServer := &http.Server{
		Handler: http.FileServer(http.Dir("/")),
	}
	go fileServer.Serve(fsListener)
	return fsListener
}
