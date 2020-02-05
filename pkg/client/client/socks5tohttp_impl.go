// +build android linux windows darwin
// +build !freebsd
// +build !openbsd
// +build !netbsd

package client

import (
	"log"

	"github.com/txthinking/brook"
	"k0s.io/k0s/pkg"
)

func (cl *client) runSocks5ToHTTP() error {
	var (
		socks5Addr = cl.GetSocks()
		httpAddr   = cl.GetSocks5ToHTTP()
	)
	if socks5Addr == "" {
		socks5Addr = pkg.SOCKS5_PROXY_PORT
	}
	log.Println("socks5tohttp:", socks5Addr, "<->", httpAddr)
	s, err := brook.NewSocks5ToHTTP(httpAddr, socks5Addr, 0, 0)
	if err != nil {
		log.Println(err)
		return err
	}
	return s.ListenAndServe()
}
