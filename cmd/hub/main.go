package main

import (
	"log"
	"os"

	//"github.com/davecgh/go-spew/spew"

	"github.com/btwiuse/conntroll/pkg/hub/hub"
	"github.com/btwiuse/conntroll/pkg/hub/config"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	c := config.Parse(os.Args[1:])

	log.Println("server is listening on", c.Port())

	log.Fatalln(hub.NewHub(c).ListenAndServe())
}
