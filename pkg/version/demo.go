//go:build version
// +build version

package main

import (
	"fmt"

	"k0s.io/pkg/version"
)

func main() {
	fmt.Print(version.Version.JsonString())
	fmt.Print(version.Version.YAMLString())
}
