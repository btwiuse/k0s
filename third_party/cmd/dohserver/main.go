package main

import (
        "log"
        "os"

	"k0s.io/third_party/pkg/cli/dohserver"
)

func main() {
        if err := dohserver.Run(os.Args[1:]); err != nil {
                log.Fatalln(err)
        }
}
