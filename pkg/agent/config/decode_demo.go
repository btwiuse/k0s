//go:build decode

package main

import (
	"fmt"
	"log"

	"k0s.io/pkg/agent/config"
)

func main() {
	conf, err := config.Decode([]byte(data))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(conf)
}

var data string = `{
  "id": "5474685a-7e87-433a-a2c8-08de1d826b42",
  "name": "modest_euler",
  "tags": [],
  "auth": "08dac2d750d4fab218dad9e586451b845f9857576f2a6f8cda214aab49765002",
  "info": {
    "os": "linux",
    "pwd": "/home/aaron/go/src/k0s.io/pkg/agent/config",
    "username": "aaron",
    "hostname": "blackarcher"
  }
}`

// "distro": "arch",
// "arch": "amd64",
