package utils

import "os"

func EnvPORT(p string) string {
	if port, ok := os.LookupEnv("PORT"); ok {
		return ":" + port
	}
	return p
}
