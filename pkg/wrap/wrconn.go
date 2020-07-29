package wrap

import (
    "net"
    "net/http"
)

func Wrconn(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
    return wrconn(w, r)
}
