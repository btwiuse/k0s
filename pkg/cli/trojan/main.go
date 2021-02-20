package trojan

import (
	"flag"
	"os"

	"github.com/p4gefau1t/trojan-go/option"

	_ "github.com/p4gefau1t/trojan-go/component"
)

func Run(args []string) error {
	os.Args = append([]string{"trojan"}, args...)
	flag.Parse()
	for {
		h, err := option.PopOptionHandler()
		if err != nil {
			return err
		}
		err = h.Handle()
		if err == nil {
			break
		}
	}
	return nil
}
