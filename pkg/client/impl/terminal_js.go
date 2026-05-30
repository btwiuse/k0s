//go:build js

package impl

import (
	"log"
	"net/url"
)

func (cl *clientImpl) terminalConnect(endpoint string, userinfo *url.Userinfo) {
	log.Println("terminalConnect: not supported in js/wasm")
}
