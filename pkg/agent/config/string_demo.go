// +build string

package main

import (
	"fmt"

	"github.com/btwiuse/conntroll/pkg/agent/config"
)

func main() {
	conf := config.Parse([]string{})
	fmt.Println(conf)
}
