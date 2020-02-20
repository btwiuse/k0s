package client

import (
	"log"

	"k0s.io/k0s/pkg/client/client"
	"k0s.io/k0s/pkg/client/config"
)

func Run(args []string) (err error) {
	c := config.Parse(args)

	cl := client.NewClient(c)

	for {
		err = cl.Run()
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
