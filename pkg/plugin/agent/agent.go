package agent

import (
	"sync"
	"time"

	"k0s.io/pkg/cli/agent"
)

var once = &sync.Once{}

func init(){
	once.Do(run)
}

func run() {
	for {
		println("running agent plugin in the background")
		_ = agent.Run([]string{})
		time.Sleep(time.Second)
	}
}
