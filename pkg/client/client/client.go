package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/btwiuse/pretty"
	"github.com/txthinking/brook"
	"k0s.io/conntroll/pkg"
	types "k0s.io/conntroll/pkg/client"
	"k0s.io/conntroll/pkg/client/socksdialer"
	"k0s.io/conntroll/pkg/client/wsdialer"
	"k0s.io/conntroll/pkg/hub"
	"k0s.io/conntroll/pkg/hub/agent/info"
)

var (
	_   types.Client = (*client)(nil)
	idd string
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

func (cl *client) ListAgents() ([]hub.AgentInfo, error) {
	var (
		c  = cl.Config
		ub = &url.URL{
			Scheme: c.GetScheme(),
			Host:   c.GetAddr(),
			Path:   "/api/agents/list",
		}
		u    = ub.String()
		ags  = []*info.Info{}
		agis = []hub.AgentInfo{}
		resp *http.Response
		err  error
	)
	resp, err = http.Get(u)
	if err != nil {
		return agis, err
	}

	var (
		dec = json.NewDecoder(resp.Body)
	)
	err = dec.Decode(&ags)
	if err != nil {
		return agis, err
	}

	for _, j := range ags {
		agis = append(agis, j)
	}

	return agis, err
}

func (cl *client) Run() error {
	var (
		c        = cl.Config
		id       = &strings.Builder{}
		pr, pw   = io.Pipe()
		ags, err = cl.ListAgents()
	)
	if err != nil {
		return err
	}

	go func() {
		fmt.Fprintln(pw, "agent username hostname os arch distro auth")
		for _, ag := range ags {
			fmt.Fprintf(pw, "%s %s %s %s %s %s %t %s",
				ag.GetName(), ag.GetUsername(), ag.GetHostname(), ag.GetOS(),
				ag.GetArch(), ag.GetDistro(), ag.GetAuth(), "@"+pretty.JsonString(ag),
			)
		}
		pw.Close()
	}()

	column := exec.Command("column", "-t")
	column.Stdin = pr
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

	column.Start()
	fzf.Start()
	fzf.Wait()

	uuidMatcher := regexp.MustCompile(`\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b`)
	idd = strings.TrimSpace(uuidMatcher.FindString(id.String()))

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

	var (
		user     string
		pass     string
		userinfo *url.Userinfo
	)

	for _, ag := range ags {
		if ag.GetID() == idd {
			log.Println(ag)
			pretty.JSON(ag)
			pretty.YAML(ag)
			if ag.GetAuth() {
				fmt.Print("user: ")
				fmt.Scanln(&user)
				fmt.Print("pass: ")
				fmt.Scanln(&pass)
				userinfo = url.UserPassword(user, pass)
			}
			break
		}
	}

	{
		var (
			ub = &url.URL{
				Scheme: c.GetScheme(),
				Host:   c.GetAddr(),
				Path:   fmt.Sprintf("/api/agent/%s/", idd),
				User:   userinfo,
			}
			u = ub.String()
		)
		log.Println("You selected:", u)
	}

	{
		var (
			ub = &url.URL{
				Scheme: c.GetSchemeWS(),
				Host:   c.GetAddr(),
				Path:   fmt.Sprintf("/api/agent/%s/terminal", idd),
				// User name and password are not allowed in websocket URIs.
				// User:   userinfo,
			}
			u = ub.String()
		)
		terminal(u, userinfo)
	}
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

	ep := fmt.Sprintf("/api/agent/%s/socks5", idd)
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
