package agent

import (
	"sync"
	"time"

	"k0s.io/pkg"
	"k0s.io/pkg/agent/agent"
	"k0s.io/pkg/agent/config"
	"k0s.io/pkg/uuid"
	"k0s.io/pkg/version"
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
		ID:      uuid.New(),
		Hub:     pkg.DEFAULT_HUB_ADDRESS,
		Tags:    []string{},
		Version: version.GetVersion(),
		Name:    "embedded",
	}
	ag := agent.NewAgent(c)

	for {
		println("running agent plugin in the background")
		// _ = agent.Run([]string{"-name", "embedded-client", "-c", "/dev/null", "-tags", "embedded-client"})

		ag.ConnectAndServe()

		time.Sleep(time.Second)
	}
}
