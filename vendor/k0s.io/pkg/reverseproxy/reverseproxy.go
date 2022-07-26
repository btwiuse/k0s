package reverseproxy

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
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
		Director: func(r *http.Request) {
			// log.Println("request as received by reverse proxy:")
			// dumpRequest(r)

			r.URL.Scheme = rpURL.Scheme
			r.Host = rpURL.Host
			r.URL.Host = r.Host

			// log.Println("request after modification by reverse proxy:")
			// dumpRequest(r)
		},
		ModifyResponse: func(r *http.Response) (err error) {
			// log.Println("reverse proxy response status code:", r.StatusCode)
			for r.StatusCode == 301 || r.StatusCode == 302 {
				if err = resolve(r); err != nil {
					return err
				}
			}
			return nil
		},
	}
	return rp
}

func resolve(r *http.Response) error {
	resp, err := http.Get(r.Header.Get("Location"))
	if err != nil {
		return err
	}
	r.Status = resp.Status
	r.StatusCode = resp.StatusCode
	r.Proto = resp.Proto
	r.ProtoMajor = resp.ProtoMajor
	r.ProtoMinor = resp.ProtoMinor
	r.Header = resp.Header
	r.Body = resp.Body
	r.ContentLength = resp.ContentLength
	r.TransferEncoding = resp.TransferEncoding
	r.Close = resp.Close
	r.Uncompressed = resp.Uncompressed
	r.Trailer = resp.Trailer
	r.Request = resp.Request
	r.TLS = resp.TLS
	// log.Println("redirected reverse proxy response status code:", r.StatusCode)
	return nil
}
