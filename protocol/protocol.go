package protocol

import (
	"errors"
	"log"
	"os/exec"
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

type Bash struct{}

func (h *Bash) Execute(req Request, res *Response) error {
	log.Println("Bash called with", req)

	if req.Command == "" {
		return errors.New("command cannot be empty")
	}
	cmd := exec.Command("/bin/bash", "-c", req.Command)
	out, err := cmd.CombinedOutput()
	res.Message = string(out)
	return err
}
