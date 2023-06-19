package reverseproxy

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"k0s.io/pkg/log"
)

func dumpRequest(r *http.Request) {
	r.Clone(context.Background()).Write(os.Stderr)
}

func Handler(addr string) http.Handler {
	rpURL, err := url.Parse(addr)
	if err != nil {
		log.Fatalln(err)
	}
	rp := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetURL(rpURL)
			r.SetXForwarded()
		},
	}
	return rp
}
