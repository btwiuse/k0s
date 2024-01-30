package version

import (
	"encoding/json"
	"fmt"

	"github.com/btwiuse/version"
)

func Run(args []string) error {
	v := version.Info
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
	return nil
}
