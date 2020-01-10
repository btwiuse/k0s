// +build !gorilla !nhooyr

package hub

import (
	"net"
	"net/http"

	"k0s.io/conntroll/pkg/wrap"
)

func wrconn(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
	return wrap.Hijack(w)
}
