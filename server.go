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
	scanner := bufio.NewScanner(os.Stdin)
	go io.Copy(os.Stdout, conn)
	for scanner.Scan() {
		line := scanner.Text()
		conn.Write([]byte(line))
	}
}

func main() {
	http.HandleFunc("/", hijacker)
	log.Println("listening on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
