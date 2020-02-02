// +build version

package main

import (
	"fmt"

	"k0s.io/k0s/pkg/version"
)

func main() {
	fmt.Print(version.Version.JsonString())
	fmt.Print(version.Version.YAMLString())
}
