// Package wetty provides a protocol and an implementation to
// control terminals thorough networks.
package wetty

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Protocols defines the name of this protocol,
// which is supposed to be used to the subprotocol of Websockt streams.
// https://stackoverflow.com/questions/37984320/why-doesnt-golang-allow-const-maps
var (
	Protocols = []string{
		"wetty",
	}
	Dialer = &websocket.Dialer{
		Subprotocols:      Protocols,
		EnableCompression: true,
	}
	Upgrader = &websocket.Upgrader{
		Subprotocols:      Protocols,
		EnableCompression: true,
		CheckOrigin:       func(r *http.Request) bool { return true },
	}
)
