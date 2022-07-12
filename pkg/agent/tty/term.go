package tty

import "fmt"

func map2arr(m map[string]string) (a []string) {
	for k, v := range m {
		a = append(a, fmt.Sprintf("%s=%s", k, v))
	}
	return
}
