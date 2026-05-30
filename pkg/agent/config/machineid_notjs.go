//go:build !js

package config

import "github.com/denisbrodbeck/machineid"

func getMachineID() (string, error) {
	return machineid.ID()
}
