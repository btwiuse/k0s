package impl

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/btwiuse/conntroll/pkg/api"
	"github.com/btwiuse/conntroll/pkg/msg"
	"github.com/btwiuse/wetty/pkg/localcmd"
	"github.com/kr/pty"
)

type Session struct {
	*localcmd.Factory
	Name string
	// client id/index, to distinguish logs of different commands
}

func (session *Session) Send(sendServer api.Session_SendServer) error {
	lc, err := session.Factory.New()
	if err != nil {
		return err
	}

	// send
	go func() {
		buf := make([]byte, 1<<12-1)
		if err != nil {
			return // err
		}
		for {
			n, err := lc.Read(buf)
			if err == io.EOF {
				return // nil
			}
			if err != nil {
				return // err
			}
			outputMsg := &api.Message{Type: msg.Type_SESSION_OUTPUT, Body: buf[:n]}
			err = sendServer.Send(outputMsg)
			if err != nil {
				return // err
			}
		}
		return // nil
	}()

	// recv
	for {
		resp, err := sendServer.Recv()
		if err != nil {
			return nil
		}
		log.Println(session.Name, resp.Type, fmt.Sprintf("%q", string(resp.Body)))
		switch resp.Type {
		case msg.Type_CLIENT_INPUT:
			_, err = lc.Write(resp.Body)
			if err != nil {
				log.Println(session.Name, "error writing to lc:", err)
				return err
			}
		case msg.Type_SESSION_RESIZE:
			sz := &pty.Winsize{}
			err = json.Unmarshal(resp.Body, sz)
			if err != nil {
				return err
			}
			err = lc.ResizeTerminal(sz)
			if err != nil {
				return err
			}
		case msg.Type_SESSION_CLOSE:
			return lc.Close()
		}
	}

	return nil
}
