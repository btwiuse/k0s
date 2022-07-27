//go:build version

package main

import (
	"encoding/json"
	"fmt"

	"k0s.io/pkg/version"
)

// go install -tags runtime_debug_buildinfo demo.go
// go install -tags runtime_debug_buildinfo k0s.io@39bfa1809a28993152c942d322b7a2c9d7ff2520

func main() {
	v := version.Info
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}
