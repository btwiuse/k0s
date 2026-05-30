//go:build js

package config

import "errors"

func getMachineID() (string, error) {
	return "", errors.New("machineid: not supported in js/wasm")
}
