package self

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/btwiuse/rng"
	"github.com/btwiuse/version"
	"k0s.io/pkg/agent/config"
	"k0s.io/pkg/agent/info"
	"k0s.io/pkg/agent/server"
)

func Agent(hub string) {
	_ = rng.NewUUID()
	c := &config.Config{
		ID:  "undefined", //rng.NewUUID(),
		Hub: hub,
		Htpasswd: map[string]string{
			"aaron": "$2a$10$WbZm/thAZI/f/QrcJn6V4OS.I61V2cLnOV.z7uXxtjHY8tZkTacLm",
		},
		Tags: []string{
			"os.Args = " + strings.Join(os.Args, " "),
			// "os.Env = " + strings.Join(os.Environ(), ":"),
		},
		Version: version.Info,
		Name:    "self",
		Info:    info.CollectInfo(),
	}

	// println(c.GetHost())
	c.WithURI()
	// println(c.GetHost())

	ag := server.NewAgent(c)

	for {
		time.Sleep(time.Second)

		log.Println(fmt.Sprintf("connecting self agent to %s", hub))
		// _ = agent.Run([]string{"-name", "embedded-client", "-c", "/dev/null", "-tags", "embedded-client"})

		log.Println(ag.ConnectAndServe())
	}
}
