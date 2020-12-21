package gitd

import (
	"log"

	"k0s.io/k0s/pkg/tunnel/listener"
	"k0s.io/k0s/pkg/gitd"
)

func Run(args []string) (err error) {
	var (
		port = args[0]
		ln   = listener.Listener(port, "/")
	)

	log.Println("server is listening on", port)

	return gitd.Serve(ln)
}
