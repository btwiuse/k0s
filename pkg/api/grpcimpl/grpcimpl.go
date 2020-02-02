package grpcimpl

import (

	// "github.com/btwiuse/wetty/pkg/msg"
	"k0s.io/k0s/pkg/agent"
	"k0s.io/k0s/pkg/api"
)

type Session struct {
	ReadOnly   bool
	TtyFactory agent.TtyFactory
	// client id/index, to distinguish logs of different commands
}

func (session *Session) Send(sendServer api.Session_SendServer) error {
	return nil
}
