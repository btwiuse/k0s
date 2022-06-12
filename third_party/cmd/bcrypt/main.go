package main

import (
	"log"
	"os"

	"github.com/btwiuse/bcrypt/pkg/bcrypt"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if err := bcrypt.Run(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
}
