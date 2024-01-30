package self

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"k0s.io/pkg/agent/config"
	"k0s.io/pkg/agent/info"
	"k0s.io/pkg/agent/server"
	"k0s.io/pkg/uuid"
	"github.com/btwiuse/version"
)

func Agent(hub string) {
	_ = uuid.New()
	c := &config.Config{
		ID:  "undefined", //uuid.New(),
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
	config.SetURI()(c)
	// println(c.GetHost())

	ag := server.NewAgent(c)

	for {
		time.Sleep(time.Second)

		log.Println(fmt.Sprintf("connecting self agent to %s", hub))
		// _ = agent.Run([]string{"-name", "embedded-client", "-c", "/dev/null", "-tags", "embedded-client"})

		log.Println(ag.ConnectAndServe())
	}
}
