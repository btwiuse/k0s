package client

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
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

	jq := exec.Command("jq", "-r", ".[].id")
	jq.Stdin = resp.Body
	// jq.Stdout = pw
	jq.Stderr = os.Stderr
	jqStdoutPipe, _ := jq.StdoutPipe()

	fzf := exec.Command("fzf")
	// fzf.Stdin = pr
	fzf.Stdin = jqStdoutPipe
	fzf.Stdout = id //os.Stdout
	fzf.Stderr = os.Stderr

	jq.Start()
	fzf.Start()
	fzf.Wait()

	wsScheme := "ws"
	if c.GetScheme() == "https" {
		wsScheme = "wss"
	}

	websocat := exec.Command("websocat",
		"tcp-l:0.0.0.0"+pkg.SOCKS5_PROXY_PORT,
		"--binary",
		fmt.Sprintf("%s://%s/api/agent/%s/socks5", wsScheme, c.GetHost(), strings.TrimSpace(id.String())),
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
	return brook.Run()
}
