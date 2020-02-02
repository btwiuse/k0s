package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/VojtechVitek/yaml-cli/pkg/cli"
	"github.com/btwiuse/pretty"
	"github.com/containerd/console"
	"github.com/txthinking/brook"
	"golang.org/x/crypto/ssh/terminal"
	"k0s.io/k0s/pkg"
	types "k0s.io/k0s/pkg/client"
	"k0s.io/k0s/pkg/client/socksdialer"
	"k0s.io/k0s/pkg/client/wsdialer"
	"k0s.io/k0s/pkg/fzf"
	"k0s.io/k0s/pkg/hub"
	"k0s.io/k0s/pkg/hub/agent/info"
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
	userinfo *url.Userinfo
}

func (cl *client) ListAgents() (agis []hub.AgentInfo, err error) {
	var (
		c  = cl.Config
		ub = &url.URL{
			Scheme: c.GetScheme(),
			Host:   c.GetAddr(),
			Path:   "/api/agents/list",
		}
		ags  = []*info.Info{}
		resp *http.Response
		req  = &http.Request{
			Method: http.MethodGet,
			URL:    ub,
		}
		t = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: c.GetInsecure(),
				},
			},
		}
	)
	resp, err = t.Do(req)
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
		w := new(tabwriter.Writer)
		w.Init(pw, 2, 0, 2, ' ', 0)
		fmt.Fprintln(w, strings.ReplaceAll("agent username hostname os arch distro auth @", " ", "\t"))
		for _, ag := range ags {
			col := fmt.Sprintf(
				strings.ReplaceAll("%s %s %s %s %s %s %t %s", " ", "\t"),
				ag.GetName(), ag.GetUsername(), ag.GetHostname(), ag.GetOS(),
				ag.GetArch(), ag.GetDistro(), ag.GetAuth(), "@"+ag.GetID(),
			)
			// log.Println(col)
			fmt.Fprintln(w, col)
		}
		w.Flush()
		pw.Close()
	}()

	/*
		(echo 'agent username hostname os arch distro auth';
		curl -s https://hub.k0s.io/api/agents/list |
		jq -cr '.[]|"\(.name) \(.username) \(.hostname) \(.os) \(.arch) \(.distro) \(.auth) @\(.)"')
		| column -t | fzf --preview 'echo {} | cut -d "@" -f 2- |jq -r .|yj -jy'
		--reverse --tac --cycle -d '@' --with-nth=1 --header-lines=1 --preview-window=right:40%
	*/

	opts := []fzf.Opt{
		fzf.WithReader(pr),
		fzf.WithWriter(id),
	}
	args := []string{
		"--tac",
		"--cycle",
		"-d", "@",
		"--reverse",
		"--with-nth", "1",
		"--header-lines", "1",
	}
	fzf.Run(fzf.ParseOptions(args, opts...), "revision")

	if strings.TrimSpace(id.String()) == "" {
		log.Fatalln("fzf empty result", id, idd)
	}

	parts := strings.Split(id.String(), "@")
	if len(parts) == 0 {
		log.Fatalln("fzf bad output format")
	}

	idd = strings.TrimSpace(parts[len(parts)-1])

	if idd == "" {
		log.Fatalln("fzf bad result", id, idd)
	}

	var (
		user string
		pass string
	)

	for _, ag := range ags {
		if ag.GetID() == idd {
			pretty.YAML(ag)
			if ag.GetAuth() {
				var (
					name  = ag.GetName()
					creds = c.GetCredentials()
				)
				ks, ok := creds[name]
				if ok && len(ks) > 0 {
					log.Println("loading cached credentials from", c.GetConfigLocation())
					// load first entry in map
					for k, v := range ks {
						user = k
						pass = v
						break
					}
				} else {
					// setup terminal
					oldState, err := terminal.MakeRaw(0)
					if err != nil {
						return err
					}

					console := console.Current()
					term := terminal.NewTerminal(console, "> ")
					term.SetPrompt("Please enter username: ")
					user, err = term.ReadLine()
					if err != nil {
						return err
					}
					pass, err = term.ReadPassword("Password: ")
					if err != nil {
						return err
					}
					terminal.Restore(0, oldState)

					{
						set := func(conf string, kv string) {
							log.Println("saving password to", conf)
							args := []string{"set", kv}
							fr, err := os.Open(conf)
							if err != nil {
								log.Println(err)
								return
							}
							defer fr.Close()

							b := bytes.NewBuffer(nil)
							err = cli.Run(b, fr, args)
							if err != nil {
								log.Println(err)
								return
							}

							err = ioutil.WriteFile(conf, b.Bytes(), 0600)
							if err != nil {
								log.Println(err)
								return
							}
						}

						if conf := c.GetConfigLocation(); conf != "" && c.GetCacheCredentials() {
							// log.Println("yes")
							set(conf, fmt.Sprintf("credentials.%s.%s: %s", ag.GetName(), user, pass))
						}
					}
				}
				cl.userinfo = url.UserPassword(user, pass)
			}
			break
		}
	}

	if len(cl.GetRedir()) > 0 {
		go cl.RunRedir()
	}
	if len(cl.GetSocks()) > 0 || len(cl.GetSocks5ToHTTP()) > 0 {
		go cl.RunSocks()
	}
	if len(cl.GetSocks5ToHTTP()) > 0 {
		go cl.RunSocks5ToHTTP()
	}

	{
		var (
			ub = &url.URL{
				Scheme: c.GetScheme(),
				Host:   c.GetAddr(),
				Path:   fmt.Sprintf("/api/agent/%s/", idd),
				User:   cl.userinfo,
			}
			u = ub.String()
		)
		log.Println("You selected:", u)
	}

	{
		cl.terminalConnect(fmt.Sprintf("/api/agent/%s/terminal", idd), cl.userinfo)
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
				return di.Dial(addr, cl.userinfo)
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
			wsconn, err := wsd.Dial(ep, cl.userinfo)
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
