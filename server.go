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
	println("start scanning")
	go io.Copy(os.Stdout, conn)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text() + "\n"
		conn.Write([]byte(line))
	}
}

func main() {
	http.HandleFunc("/", hijacker)
	log.Println("listening on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
