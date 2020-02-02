// +build string

package main

import (
	"fmt"

	"k0s.io/k0s/pkg/agent/config"
)

func main() {
	conf := config.Parse([]string{})
	fmt.Println(conf)
}
