package version

import (
	"encoding/json"
	"fmt"

	"k0s.io/pkg/version"
)

func Run(args []string) error {
	v := version.Info
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
	return nil
}
