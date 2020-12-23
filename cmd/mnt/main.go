package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

    "k0s.io/k0s/pkg/tunnel/listener"
    "k0s.io/k0s/pkg/reverseproxy"
)

func main() {
	addr := "" // "127.0.0.1" + port.Port + "/"
	from := "https://baidu.com"

	if len(os.Args) <= 1 {
		log.Fatalln("wrong number of args.")
	}

	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	if len(os.Args) > 2 {
		from = os.Args[2]
	}

	log.Println("Listening on", addr)

	// from url
	if strings.HasPrefix(from, "http") {
		log.Println(http.Serve(listener.Listener(addr, from), reverseproxy.Handler(from)))
	} else {
		// from local file/dir
		// relative file: base + filename
		// relative dir: base
		// absolute file: base
		// absolute dir: base
		dir, base := filepath.Dir(strings.TrimSuffix(from, "/")), filepath.Base(from)
		path := filepath.Clean("/" + base)
		if isDir(from) && path != "/" {
			path += "/"
		}
		log.Println(dir, base)
		log.Println(isDir(from), path)
		log.Println(http.Serve(listener.Listener(addr, path), http.FileServer(http.Dir(dir))))
	}
}

func isDir(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		log.Fatalln(err)
		return false
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true
	default:
		return false
	}
}
