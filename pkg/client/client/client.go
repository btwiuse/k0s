package client

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/txthinking/brook"
	"k0s.io/conntroll/pkg"
	types "k0s.io/conntroll/pkg/client"
	"k0s.io/conntroll/pkg/client/socksdialer"
	"k0s.io/conntroll/pkg/client/wsdialer"
)

var idd string

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

	uuidMatcher := regexp.MustCompile(`\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b`)
	idd = uuidMatcher.FindString(id.String())

	if len(cl.GetRedir()) > 0 {
		go cl.RunRedir()
	}
	if len(cl.GetSocks()) > 0 || len(cl.GetSocks5ToHTTP()) > 0 {
		go cl.RunSocks()
	}
	if len(cl.GetSocks5ToHTTP()) > 0 {
		go cl.RunSocks5ToHTTP()
	}

	if idd == "" {
		os.Exit(0)
	}
	log.Println("You selected:", fmt.Sprintf("%s://%s/api/agent/%s/", c.GetScheme(), c.GetAddr(), idd))

	wss := "wss"
	if c.GetScheme() == "http" {
		wss = "ws"
	}

	ep := fmt.Sprintf("%s://%s/api/agent/%s/terminal", wss, c.GetAddr(), idd)
	terminal(ep)
	return nil
}

func (cl *client) RunRedir() error {
	log.Println("redir", cl.GetRedir())
	return nil
}

func (cl *client) RunSocks() error {
	var (
		c = cl.Config
	)
	wsd := wsdialer.New(c)

	ep := fmt.Sprintf("/api/agent/%s/socks5", strings.TrimSpace(idd))
	log.Println("dial", ep)

	addr := pkg.SOCKS5_PROXY_PORT
	if cl.GetSocks() != "" {
		addr = cl.GetSocks()
	}
	log.Println("socks5 listening on", addr)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	// test socksdialer functionality
	// currently broken
	// need further investigation
	go func() {
		di := socksdialer.New(c, ep)
		tr := &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return di.Dial(addr)
			},
		}
		client := http.Client{Transport: tr}
		resp, err := client.Get("https://ip.sb")
		if err != nil {
			log.Println(err)
			return
		}
		io.Copy(os.Stderr, resp.Body)
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			wsconn, err := wsd.Dial(ep)
			if err != nil {
				log.Println(err)
				return
			}
			go io.Copy(wsconn, conn)
			io.Copy(conn, wsconn)
		}()
	}
	return nil
}

func (cl *client) RunSocks5ToHTTP() error {
	var (
		socks5Addr = cl.GetSocks()
		httpAddr   = cl.GetSocks5ToHTTP()
	)
	if socks5Addr == "" {
		socks5Addr = pkg.SOCKS5_PROXY_PORT
	}
	log.Println("socks5tohttp:", socks5Addr, "<->", httpAddr)
	s, err := brook.NewSocks5ToHTTP(httpAddr, socks5Addr, 0, 0)
	if err != nil {
		log.Println(err)
		return err
	}
	return s.ListenAndServe()
}
