package main

import (
	"log"
	"os"

	//"github.com/davecgh/go-spew/spew"

	"github.com/btwiuse/invctrl/pkg/hub"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	addr := ":8000"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}
	log.Println("server is listening on", addr)
	log.Fatalln(hub.NewServer(addr).ListenAndServe())
}
