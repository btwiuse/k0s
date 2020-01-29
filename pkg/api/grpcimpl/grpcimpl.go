package grpcimpl

import (

	// "github.com/btwiuse/wetty/pkg/msg"
	"k0s.io/conntroll/pkg/agent"
	"k0s.io/conntroll/pkg/api"
)

type Session struct {
	ReadOnly   bool
	TtyFactory agent.TtyFactory
	// client id/index, to distinguish logs of different commands
}

func (session *Session) Send(sendServer api.Session_SendServer) error {
	return nil
}
