package impl

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"

	"github.com/VojtechVitek/yaml-cli/pkg/cli"
	"github.com/abiosoft/ishell/v2"
	"github.com/containerd/console"
	fzf "github.com/junegunn/fzf/src"
	"golang.org/x/crypto/ssh/terminal"
	"k0s.io"
	"k0s.io/pkg/client"
	"k0s.io/pkg/client/config"
	"k0s.io/pkg/hub/agent/info"
)

var (
	_   client.Client = (*clientImpl)(nil)
	idd string
)

func NewClient(c *config.Config) client.Client {
	cl := &clientImpl{
		config: c,
		sl:     ishell.New(),
		dialer: &dialer{c},
	}
	return cl
}

type clientImpl struct {
	*dialer
	config   *config.Config
	userinfo *url.Userinfo
	sl       *ishell.Shell
}

func (cl *clientImpl) Config() *config.Config {
	return cl.config
}

func (cl *clientImpl) ListAgents() (agis []*info.Info, err error) {
	var (
		c  = cl.config
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
					InsecureSkipVerify: c.Insecure,
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

func (cl *clientImpl) MiniRun() error {
	id := os.Getenv("ID")
	endpoint := fmt.Sprintf("/api/agent/%s/terminal", id)
	log.Println(os.Args, endpoint)
	cl.terminalConnect(endpoint, nil)
	return nil
}

func (cl *clientImpl) Run() error {
	cl.sl.AddCmd(&ishell.Cmd{
		Name: "self",
		Help: "run /proc/self/exe with args",
		Func: func(c *ishell.Context) {
			self, err := os.Executable()
			if err != nil {
				log.Println(err)
				return
			}
			cmd := exec.Command(self, c.RawArgs[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr
			cmd.Run()
		},
	})
	cl.sl.AddCmd(&ishell.Cmd{
		Name: "kill",
		Help: "kill background job by job id",
		Func: func(c *ishell.Context) {
		},
	})
	cl.sl.AddCmd(&ishell.Cmd{
		Name: "jobs",
		Help: "list background jobs TODO",
		Func: func(c *ishell.Context) {
		},
	})
	cl.sl.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "login to agent with id",
		Func: func(c *ishell.Context) {
			if len(c.RawArgs) < 2 {
				log.Println("wrong number of args")
				return
			}
			id := c.RawArgs[1]
			if id == "" {
				return
			}
			cl.runLogin(id)
		},
	})
	cl.sl.AddCmd(&ishell.Cmd{
		Name: "fzf",
		Help: "list currently connected agents and pipe to fzf",
		Func: func(c *ishell.Context) {
			id := cl.runFzf()
			if id == "" {
				return
			}
			cl.runLogin(id)
		},
	})
	cl.sl.AddCmd(&ishell.Cmd{
		Name: "ls",
		Help: "list currently connected agents",
		Func: func(c *ishell.Context) {
			cl.printAgentTable(os.Stdout)
		},
	})
	cl.sl.Run()
	return nil
}

func (cl *clientImpl) printAgentTable(out io.Writer) error {
	ags, err := cl.ListAgents()
	if err != nil {
		return err
	}
	w := new(tabwriter.Writer)
	w.Init(out, 2, 0, 2, ' ', 0)
	fmt.Fprintln(w, strings.ReplaceAll("agent username hostname os arch auth @", " ", "\t"))
	for _, ag := range ags {
		col := fmt.Sprintf(
			strings.ReplaceAll("%s %s %s %s %s %t %s", " ", "\t"),
			ag.Name, ag.Username, ag.Hostname, ag.OS,
			ag.Arch, ag.Auth, "@"+ag.ID,
		)
		fmt.Fprintln(w, col)
	}
	w.Flush()
	return nil
}

// convert io.Reader to chan string using bufio.Scanner
func readerToChan(r io.Reader) chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()
	return ch
}

func (cl *clientImpl) runFzf() string {
	var (
		out    = make(chan string, 1)
		pr, pw = io.Pipe()
	)

	go func() {
		cl.printAgentTable(pw)
		pw.Close()
	}()

	/*
		(echo 'agent username hostname os arch distro auth';
		curl -s https://hub.k0s.io/api/agents/list |
		jq -cr '.[]|"\(.name) \(.username) \(.hostname) \(.os) \(.arch) \(.distro) \(.auth) @\(.)"')
		| column -t | fzf --preview 'echo {} | cut -d "@" -f 2- |jq -r .|yj -jy'
		--reverse --tac --cycle -d '@' --with-nth=1 --header-lines=1 --preview-window=right:40%
	*/

	args := []string{
		"--tac",
		"--cycle",
		"-d", "@",
		"--reverse",
		"--with-nth", "1",
		"--header-lines", "1",
	}

	fzfApp, err := fzf.ParseOptions(true, args)
	if err != nil {
		log.Fatalln(err)
	}

	fzfApp.Input = readerToChan(pr)
	fzfApp.Output = out

	exitCode, err := fzf.Run(fzfApp)
	if exitCode != 0 || err != nil {
		return ""
	}

	id := <-out

	if strings.TrimSpace(id) == "" {
		log.Fatalln("fzf empty result", id, idd)
		// return nil
	}

	parts := strings.Split(id, "@")
	if len(parts) == 0 {
		log.Fatalln("fzf bad output format")
	}

	idd = strings.TrimSpace(parts[len(parts)-1])

	if idd == "" {
		log.Fatalln("fzf bad result", id, idd)
	}
	return idd
}

func (cl *clientImpl) runLogin(idd string) error {
	var (
		c    = cl.config
		user string
		pass string
	)

	ags, err := cl.ListAgents()
	if err != nil {
		return err
	}
	for _, ag := range ags {
		if ag.ID == idd {
			if *ag.Auth {
				var (
					name  = ag.Name
					creds = c.Credentials
				)
				ks, ok := creds[name]
				if ok && len(ks) > 0 {
					log.Println("loading cached credentials from", c.ConfigLocation)
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

						if conf := c.ConfigLocation; conf != "" && c.CacheCredentials {
							// log.Println("yes")
							set(conf, fmt.Sprintf("credentials.%s.%s: %s", ag.Name, user, pass))
						}
					}
				}
				cl.userinfo = url.UserPassword(user, pass)
			}
			break
		}
	}

	if len(cl.config.Redir) > 0 {
		go cl.RunRedir()
	}
	if len(cl.config.Socks) > 0 {
		go cl.RunSocks()
	}
	if len(cl.config.Doh) > 0 {
		go cl.RunDoh()
	}

	cl.terminalConnect(fmt.Sprintf("/api/agent/%s/terminal", idd), cl.userinfo)
	return nil
}

func (cl *clientImpl) RunRedir() error {
	ep := fmt.Sprintf("/api/agent/%s/redir", idd)
	log.Println("dial", ep)

	addr := k0s.REDIR_PROXY_PORT
	if cl.config.Redir != "" {
		addr = cl.config.Redir
	}
	log.Println("redir listening on", addr)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			wsconn, err := cl.Dial(ep, cl.userinfo)
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

func (cl *clientImpl) RunSocks() error {
	ep := fmt.Sprintf("/api/agent/%s/socks5", idd)
	log.Println("dial", ep)

	addr := k0s.SOCKS5_PROXY_PORT
	if cl.config.Socks != "" {
		addr = cl.config.Socks
	}
	log.Println("socks5 listening on", addr)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			wsconn, err := cl.Dial(ep, cl.userinfo)
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

func (cl *clientImpl) RunDoh() error {
	ep := fmt.Sprintf("/api/agent/%s/doh", idd)
	log.Println("dial", ep)

	addr := k0s.DOH_PROXY_PORT
	if cl.config.Doh != "" {
		addr = cl.config.Doh
	}
	log.Println("doh listening on", addr)
	ln, err := net.Listen("udp", addr)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			wsconn, err := cl.Dial(ep, cl.userinfo)
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

func authorizationHeader(userinfo *url.Userinfo) http.Header {
	return http.Header{
		"Authorization": {
			"Basic " + base64.StdEncoding.EncodeToString([]byte(userinfo.String())),
		},
	}
}
