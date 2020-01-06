// +build !gorilla !nhooyr

package hub

import (
	"net"
	"net/http"

	"github.com/btwiuse/conntroll/pkg/wrap"
)

func wrconn(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
	return wrap.Hijack(w)
}
