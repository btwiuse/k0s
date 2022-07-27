//go:build runtime_debug_buildinfo

package version

import (
	"encoding/json"
	"runtime/debug"
)

func init() {
	v, ok := debug.ReadBuildInfo()
	if ok {
		b, _ := json.MarshalIndent(v, "", "  ")
		println(string(b))
		println(v.Main.Version)
	}
}
