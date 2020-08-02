package listener

import (
	"log"
	"net"
)

func Listener(port string) net.Listener { return listener(port) }

func listener(p string) net.Listener {
	ln, err := net.Listen("tcp", p)
	if err != nil {
		log.Fatalln(err)
	}
	return ln
}
