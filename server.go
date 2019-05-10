package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

func hijacker(w http.ResponseWriter, r *http.Request) {
	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		log.Println(nil)
	}
	log.Println("connected:", conn.RemoteAddr())

	quit := make(chan struct{})
	copy := func(dst io.Writer, src io.Reader) {
		defer log.Println("disconnected:", conn.RemoteAddr())
		defer close(quit)
		buf := make([]byte, 1)
		for {
			_, err := src.Read(buf)
			if err != nil {
				return
			}
			dst.Write(buf)
		}
	}
	go copy(os.Stdout, conn)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-quit:
			return
		default:
			if scanner.Scan() {
				line := scanner.Text() + "\n"
				conn.Write([]byte(line))
			} else {
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/", hijacker)
	log.Println("listening on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
