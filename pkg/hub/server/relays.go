package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/btwiuse/wsconn"
	"github.com/gorilla/mux"
	"k0s.io/pkg/api"
	"k0s.io/pkg/hub"
)

func protocolRelay(protocol api.ProtocolID, ag hub.Agent) http.HandlerFunc {
	println("protocolRelay", protocol)
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			vars   = mux.Vars(r)
			id     = vars["id"]
			prefix = fmt.Sprintf("/api/agent/%s/%s", id, protocol)
		)
		// websocket
		if strings.ToLower(r.Header.Get("Upgrade")) == "websocket" && strings.ToLower(r.Header.Get("Connection")) == "upgrade" {
			wsc, err := wsconn.Wrconn(w, r)
			if err != nil {
				log.Println(err)
				return
			}
			defer wsc.Close()

			conn := ag.OpenChannel(protocol)
			defer conn.Close()

			// ensure both copy goroutines are done before returning
			done := make(chan struct{}, 2)
			go func() {
				_, err := io.Copy(conn, wsc)
				if err != nil {
					log.Println(err)
				}
				done <- struct{}{}
			}()

			go func() {
				_, err := io.Copy(wsc, conn)
				if err != nil {
					log.Println(err)
				}
				done <- struct{}{}
			}()

			<-done
			<-done

			return
		}
		// http
		dialCtx := func(ctx context.Context, network, addr string) (net.Conn, error) {
			conn := ag.OpenChannel(protocol)
			return conn, nil
		}
		tr := &http.Transport{
			DialContext:     dialCtx,
			MaxIdleConns:    100,
			IdleConnTimeout: 90 * time.Second,
		}
		rewrite := func(req *httputil.ProxyRequest) {
			req.SetXForwarded()
			req.Out.URL.Host = r.Host
			req.Out.URL.Scheme = "http"
		}
		rp := &httputil.ReverseProxy{
			Rewrite:   rewrite,
			Transport: tr,
		}
		http.StripPrefix(prefix, rp).ServeHTTP(w, r)
	}
}
