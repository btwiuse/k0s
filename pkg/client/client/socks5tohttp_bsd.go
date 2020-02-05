// +build freebsd openbsd netbsd

package client

import (
	"log"
)

func (cl *client) runSocks5ToHTTP() error {
	log.Println("socks5tohttp not implemented for *BSD, due to limitations of the dependent brook package")
	return nil
}
