package agent

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/btwiuse/rng"
	"github.com/btwiuse/version"
	"k0s.io"
	"k0s.io/pkg/agent/config"
	"k0s.io/pkg/agent/server"
)

var once = &sync.Once{}

func init() {
	once.Do(gorun)
}

func gorun() {
	go run()
}

func run() {
	c := &config.Config{
		ID:  rng.NewUUID(),
		Hub: k0s.DEFAULT_HUB_ADDRESS,
		Htpasswd: map[string]string{
			"aaron": "$2a$10$WbZm/thAZI/f/QrcJn6V4OS.I61V2cLnOV.z7uXxtjHY8tZkTacLm",
		},
		Tags: []string{
			"os.Args = " + strings.Join(os.Args, " "),
			// "os.Env = " + strings.Join(os.Environ(), ":"),
		},
		Version: version.Info,
		Name:    "embedded",
	}
	ag := server.NewAgent(c)

	for {
		println("running agent plugin in the background")
		// _ = agent.Run([]string{"-name", "embedded-client", "-c", "/dev/null", "-tags", "embedded-client"})

		ag.ConnectAndServe()

		time.Sleep(time.Second)
	}
}
