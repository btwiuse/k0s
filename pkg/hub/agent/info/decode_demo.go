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

	conf, err = config.Decode([]byte(data2))
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
  "meta": {
    "os": "linux",
    "pwd": "/home/aaron/go/src/k0s.io/pkg/agent/config",
    "username": "aaron",
    "hostname": "blackarcher"
  }
}`

var data2 string = `{
  "id": "f97a2b7b-1ba2-4c6e-ad2c-4569a6f8145e",
  "name": "mystifying_yalow",
  "tags": [
    "demo"
  ],
  "htpasswd": null,
  "meta": {
    "os": "linux",
    "pwd": "/home/aaron/go/src/github.com/btwiuse/k0s",
    "arch": "amd64",
    "username": "aaron",
    "hostname": "localhost"
  },
  "version": {
    "GitCommit": "ae38a8bc007abd55a03fccc32ac8eb8c3a947dc4",
    "GitState": "dirty",
    "GitBranch": "master",
    "GitSummary": "v0.0.1-78-gae38a8b-dirty",
    "BuildDate": "2020-01-17T02:16:53Z",
    "Version": "v0.0.1"
  }
}`
