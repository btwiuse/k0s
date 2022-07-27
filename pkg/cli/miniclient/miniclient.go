package miniclient

import (
	"log"
	"time"

	"k0s.io/pkg/client/config"
	"k0s.io/pkg/client/impl"
)

/*

./target/debug/hub_rs

./k0s.io agent :8000

ID=9a541eea-0a59-4af1-88e3-868f7bdec8fc ./k0s.io miniclient :8000

*/

func Run(args []string) (err error) {
	c := config.Parse(args)

	cl := impl.NewClient(c)

	for {
		err = cl.MiniRun()
		if err != nil {
			log.Println(err)
		}
		log.Println("Reconnecting in 3 Seconds")
		time.Sleep(3 * time.Second)
	}
	return nil
}
