package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"k0s.io/conntroll/pkg"
	types "k0s.io/conntroll/pkg/client"
)

func NewClient(c types.Config) types.Client {
	cl := &client{
		Config: c,
	}
	return cl
}

type client struct {
	types.Config
}

func (cl *client) Run() error {
	var (
		c = cl.Config
	)
	resp, err := http.Get(c.GetScheme() + "://" + c.GetAddr() + "/api/agents/list")
	if err != nil {
		log.Fatalln(err)
	}

	// id := &bytes.Buffer{}
	id := &strings.Builder{}

	head := exec.Command("echo", "agent username hostname os arch distro auth")
	headStdoutPipe, _ := head.StdoutPipe()

	jq := exec.Command("jq", "-cr", `.[]|"\(.name) \(.username) \(.hostname) \(.os) \(.arch) \(.distro) \(.auth) @\(.)"`)
	jq.Stdin = resp.Body
	jq.Stderr = os.Stderr
	jqStdoutPipe, _ := jq.StdoutPipe()

	column := exec.Command("column", "-t")
	column.Stdin = io.MultiReader(headStdoutPipe, jqStdoutPipe)
	columnStdoutPipe, _ := column.StdoutPipe()

	/*
		(echo 'agent username hostname os arch distro auth';
		curl -s https://hub.k0s.io/api/agents/list |
		jq -cr '.[]|"\(.name) \(.username) \(.hostname) \(.os) \(.arch) \(.distro) \(.auth) @\(.)"')
		| column -t | fzf --preview 'echo {} | cut -d "@" -f 2- |jq -r .|yj -jy'
		--reverse --tac --cycle -d '@' --with-nth=1 --header-lines=1 --preview-window=right:40%
	*/

	fzf := exec.Command("fzf",
		"--tac",
		"--cycle",
		"-d", "@",
		"--reverse",
		"--with-nth", "1",
		"--header-lines", "1",
		"--preview-window", "right:40%",
		"--preview", "echo {}|cut -d @ -f 2-|jq -r .|yj -jy",
	)
	// fzf.Stdin = pr
	fzf.Stdin = columnStdoutPipe
	fzf.Stdout = id //os.Stdout
	fzf.Stderr = os.Stderr

	head.Start()
	jq.Start()
	column.Start()
	fzf.Start()
	fzf.Wait()

	wsScheme := "ws"
	if c.GetScheme() == "https" {
		wsScheme = "wss"
	}

	uuidMatcher := regexp.MustCompile(`\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b`)
	idd := uuidMatcher.FindString(id.String())

	websocat := exec.Command("websocat",
		"tcp-l:0.0.0.0"+pkg.SOCKS5_PROXY_PORT,
		"--binary",
		fmt.Sprintf("%s://%s/api/agent/%s/socks5", wsScheme, c.GetHost(), strings.TrimSpace(idd)),
	)
	brook := exec.Command("brook",
		"socks5tohttp",
		"-l", pkg.HTTP_PROXY_PORT,
		"-s", pkg.SOCKS5_PROXY_PORT,
	)
	websocat.Stdout = os.Stdout
	websocat.Stderr = os.Stderr
	brook.Stdout = os.Stdout
	brook.Stderr = os.Stderr
	websocat.Start()

	{
		if len(cl.GetRedir()) > 0 {
			go cl.RunRedir()
		}
		if len(cl.GetSocks()) > 0 {
			go cl.RunSocks()
		}
	}
	return brook.Run()
}

func (cl *client) RunRedir() error {
	log.Println("redir", cl.GetRedir())
	return nil
}

func (cl *client) RunSocks() error {
	log.Println("socks", cl.GetSocks())
	return nil
}
