//go:build js

package tty

import (
	"errors"

	"k0s.io/pkg/agent"
)

func New(args []string) (agent.Tty, error) {
	return nil, errors.New("tty: not supported in js/wasm")
}

func NewEnv(args []string, env map[string]string) (agent.Tty, error) {
	return nil, errors.New("tty: not supported in js/wasm")
}
