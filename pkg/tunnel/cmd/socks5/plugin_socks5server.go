package main

import (
	"log"
	"net"
	"net/http"

	"github.com/btwiuse/gost"
	"k0s.io/pkg/tunnel/listener"
)

const HOST = "ws://127.0.0.1:8000/api/chassis/socks5/"

var autoHandler = gost.AutoHandler()
var teleportListener = listener.Listener(HOST, "/")
var fsHandler = http.FileServer(http.Dir("/tmp/"))

func autoServe(l net.Listener) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Println("socks5server: recovered from panic:", r, c)
				}
			}()
			autoHandler.Handle(c)
			c.Close()
		}()
	}
}

func main() {
	log.Println("listening on", HOST)
	// http.Serve(teleportListener, fsHandler)
	autoServe(teleportListener)
}
