// +build raw
// +build !gorilla
// +build !nhooyr

package wrap

import (
	"net"
	"net/http"
)

func wrconn(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
	return Hijack(w)
}
