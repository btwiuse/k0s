package gitd

import (
	"log"

	"k0s.io/k0s/pkg/gitd"
	"k0s.io/k0s/pkg/tunnel/listener"
)

func Run(args []string) (err error) {
	var (
		port = args[0]
		ln   = listener.Listener(port, "/")
	)

	log.Println("server is listening on", port)

	return gitd.Serve(ln)
}
