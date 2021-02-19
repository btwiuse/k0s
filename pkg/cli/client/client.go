package client

import (
	"log"
	"time"

	"k0s.io/pkg/client/client"
	"k0s.io/pkg/client/config"
)

func Run(args []string) (err error) {
	c := config.Parse(args)

	cl := client.NewClient(c)

	for {
		err = cl.Run()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
	}
	return nil
}
