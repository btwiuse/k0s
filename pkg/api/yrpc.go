package api

import (
	"github.com/yrpc/yrpc"
)

const (
	// caller: both
	// callee: both
	// payload: optional
	Ping yrpc.Cmd = iota
	Pong

	// caller: agent
	// callee: hub
	// payload: agent info
	AgentRegisterRequest
	AgentRegisterResponse

	// caller: hub
	// callee: agent
	// payload: optional
	AcceptRequest
	AcceptResponse
)
