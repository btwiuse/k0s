//go:build string

package main

import (
	"fmt"

	"k0s.io/pkg/agent/config"
)

func main() {
	conf := config.Parse([]string{})
	fmt.Println(conf)
}
