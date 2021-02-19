package wetty

import (
	"log"
	"strings"

	"k0s.io/pkg/wetty"
)

func Run(args []string) error {
	if len(args) == 0 {
		log.Fatalln("usage: wetty [command] [args]...")
	}

	log.Printf("WeTTY is starting with command: %s", strings.Join(args, " "))
	return wetty.New(args).Run()
}
