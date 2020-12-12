package chassis

import (
	"log"
	"os"

	"k0s.io/k0s/pkg/cli/chassis"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	chassis.Run(os.Args[1:])
}
