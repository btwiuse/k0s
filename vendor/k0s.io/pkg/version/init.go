//go:build runtime_debug_buildinfo
// +build runtime_debug_buildinfo

package version

import (
	"runtime/debug"

	"github.com/btwiuse/pretty"
)

func init() {
	b, ok := debug.ReadBuildInfo()
	if ok {
		println(pretty.YAMLString(b))
	}
}
