//go:build version
// +build version

package main

import (
	"fmt"

	"k0s.io/pkg/version"
)

func main() {
	v := &version.Version{}
	fmt.Print(v.JsonString())
	fmt.Print(v.YAMLString())
}
