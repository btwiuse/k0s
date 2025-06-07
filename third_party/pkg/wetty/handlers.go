package wetty

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"modernc.org/httpfs"

	"github.com/btwiuse/wetty/pkg/assets"
	"github.com/btwiuse/wsconn"
	asciitransport "k0s.io/third_party/pkg/asciitransport/v2"
	"k0s.io/third_party/pkg/localcmd"
	"k0s.io/third_party/pkg/wetty/wetty"
)

func (server *Server) setupHandlers(pathPrefix string) http.Handler {
	mux := http.NewServeMux()
	staticFileServer := http.FileServer(httpfs.NewFileSystem(assets.Assets, time.Now()))
	mux.Handle(pathPrefix, http.StripPrefix(pathPrefix, staticFileServer))
	mux.HandleFunc(pathPrefix+"terminal", server.terminalHandler)
	return handlers.LoggingHandler(os.Stderr, mux)
}

var term0 *localcmd.Lc

func (server *Server) terminalHandler(w http.ResponseWriter, r *http.Request) {
	closeReason := "unknown reason"

	defer func() {
		log.Printf(
			"Connection closed by %s: %s, connections: %d/%d",
			closeReason, r.RemoteAddr, 0, 0,
		)
	}()

	log.Printf("New client connected: %s, connections: %d/%d", r.RemoteAddr, 0, 0)

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed /*405*/)
		return
	}

	conn, err := wetty.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		closeReason = err.Error()
		return
	}
	defer conn.Close()
	wsc := wsconn.NetConn(conn)

	var term *localcmd.Lc
	if term0 == nil {
		term0, _ = server.factory.New()
	}
	term = term0
	var (
		opts = []asciitransport.Opt{
			asciitransport.WithReader(term),
			asciitransport.WithWriter(term),
			asciitransport.WithResizeHook(func(w, h uint16) {
				// cols, rows
				err := term.Resize(int(h), int(w))
				if err != nil {
					log.Println(err)
				}
			}),
		}
		aserver = asciitransport.Server(wsc, opts...)
	)

	<-aserver.Done()
	log.Println("detached")
	//log.Println("detached", term.Close())
}
