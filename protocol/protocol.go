package protocol

import (
	"errors"
	"log"
)

type Request struct {
	Command string
}

type Response struct {
	Message string
}

type Hello struct{}

func (h *Hello) Execute(req Request, res *Response) error {
	log.Println("Hello called with", req)
	if req.Command == "" {
		return errors.New("command cannot be empty")
	}
	res.Message = "Hello " + req.Command
	return nil
}
