package mnt

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

    "k0s.io/k0s/pkg/tunnel/listener"
    "k0s.io/k0s/pkg/reverseproxy"
)

func Run(args []string) error {
	addr := "" // "127.0.0.1" + port.Port + "/"
	from := "https://baidu.com"

	if len(args) == 0 {
		log.Fatalln("wrong number of args.")
	}

	if len(args) > 0 {
		addr = args[0]
	}

	if len(args) > 1 {
		from = args[1]
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

    return nil
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
