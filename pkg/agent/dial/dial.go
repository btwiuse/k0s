package dial

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/url"
	"os/exec"
	"strings"

	"github.com/btwiuse/invctrl/pkg/agent/config"
	"github.com/btwiuse/invctrl/pkg/header"
	"github.com/btwiuse/pretty"
	"github.com/gorilla/websocket"
)

func run(oneliner string) []byte {
	cmd := exec.Command("/bin/bash", "-c", oneliner)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return []byte(err.Error())
	}
	return out
}

var (
	pwd      = strings.TrimSpace(string(run(`pwd`)))
	whoami   = strings.TrimSpace(string(run(`whoami`)))
	hostname = strings.TrimSpace(string(run(`hostname`)))
)

func WsDial(url string) *websocket.Conn {
	dialer := &websocket.Dialer{}
	ws, _, err := dialer.Dial(url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return ws
}

func DialSlave(hostport string, args url.Values) net.Conn {
	basePath := "/new/slave"
	queryString := args.Encode()
	log.Println(queryString)
	requestText := fmt.Sprintf("GET %s?%s HTTP/1.0\r\n\r\n", basePath, queryString)

	conn, err := net.Dial("tcp", hostport)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = conn.Write([]byte(requestText))
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}

func DialAgent(hostport string) net.Conn {
	basePath := "/new/agent"
	var agentInfo url.Values = make(map[string][]string)
	agentInfo.Set("id", config.Default.Id)
	agentInfo.Set("pwd", pwd)
	agentInfo.Set("whoami", whoami)
	agentInfo.Set("hostname", hostname)
	queryString := agentInfo.Encode()
	log.Println(queryString)
	requestText := fmt.Sprintf("GET %s?%s HTTP/1.0\r\n\r\n", basePath, queryString)

	conn, err := net.Dial("tcp", hostport)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = conn.Write([]byte(requestText))
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}

func writeConnAppend(conn net.Conn, nonce string) {
	header := pretty.JSONString(&header.AgentHeader{
		Append: true,
		Id:     config.Default.Id,
		Nonce:  nonce,
	})
	conn.Write([]byte(header))
}

func writeAgentInfo(conn net.Conn) {
	header := pretty.JSONString(&header.AgentHeader{
		Pwd:      pwd,
		Whoami:   whoami,
		Hostname: hostname,
	})
	conn.Write([]byte(header))
}

func readConn(conn net.Conn) string {
	cheader := &header.SlaveHeader{}
	if err := json.NewDecoder(conn).Decode(&cheader); err != nil {
		log.Fatalln(err)
	}
	if cheader.Status != "OK" {
		log.Fatalln(cheader.Status, cheader.Err)
	}
	return cheader.Id
}

func Handshake(conn net.Conn) string {
	writeAgentInfo(conn)
	agentId := readConn(conn)
	return agentId
}

func HandshakeAppend(conn net.Conn, nonce string) {
	writeConnAppend(conn, nonce)
	readConn(conn)
}
