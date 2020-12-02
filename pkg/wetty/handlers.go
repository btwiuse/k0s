package wetty

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"modernc.org/httpfs"

	"github.com/btwiuse/wetty/pkg/assets"
	asciitransport "k0s.io/k0s/pkg/asciitransport/v2"
	"k0s.io/k0s/pkg/utils"
	"k0s.io/k0s/pkg/wetty/wetty"
    "k0s.io/k0s/pkg/localcmd"
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
	wsconn := utils.NetConn(conn)

	var	term *localcmd.Lc
    if term0 == nil {
        term0, _ = server.factory.New()
    }
    term = term0
	var (
		opts    = []asciitransport.Opt{
			asciitransport.WithReader(term),
			asciitransport.WithWriter(term),
			asciitransport.WithResizeHook(func(w, h uint16){
				// cols, rows
				err := term.Resize(int(h), int(w))
				if err != nil {
					log.Println(err)
				}
			}),
		}
		aserver = asciitransport.Server(wsconn, opts...)
	)

	<-aserver.Done()
	log.Println("detached")
	//log.Println("detached", term.Close())
}
