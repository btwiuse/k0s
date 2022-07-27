package client

import (
	"log"
	"time"

	"k0s.io/pkg/client/config"
	"k0s.io/pkg/client/impl"
)

func Run(args []string) (err error) {
	c := config.Parse(args)

	cl := impl.NewClient(c)

	for {
		err = cl.Run()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
	}
	return nil
}
