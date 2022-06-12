package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/btwiuse/dkg/pkg/imgpkg/cmd"
	"github.com/cppforlife/go-cli-ui/ui"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	log.SetOutput(ioutil.Discard)

	// TODO logs
	// TODO log flags used

	confUI := ui.NewConfUI(ui.NewNoopLogger())
	defer confUI.Flush()

	command := cmd.NewDefaultImgpkgCmd(confUI)

	err := command.Execute()
	if err != nil {
		confUI.ErrorLinef("Error: %v", err)
		os.Exit(1)
	}

	confUI.PrintLinef("Succeeded")
}
