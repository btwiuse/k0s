package impl

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/btwiuse/conntroll/pkg/api"
	"github.com/btwiuse/wetty/pkg/localcmd"
	"github.com/btwiuse/wetty/pkg/message"
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
			outputMsg := &api.Message{Type: []byte{message.SessionOutput}, Body: buf[:n]}
			err = sendServer.Send(outputMsg)
			if err != nil {
				return // err
			}
		}
		return // nil
	}()

	// recv
	for {
		msg, err := sendServer.Recv()
		if err != nil {
			return nil
		}
		log.Println(session.Name, message.ToString(msg.Type[0]), fmt.Sprintf("%q", string(msg.Body)))
		switch msgType := msg.Type[0]; msgType {
		case message.ClientInput:
			_, err = lc.Write(msg.Body)
			if err != nil {
				log.Println(session.Name, "error writing to lc:", err)
				return err
			}
		case message.SessionResize:
			sz := &pty.Winsize{}
			err = json.Unmarshal(msg.Body, sz)
			if err != nil {
				return err
			}
			err = lc.ResizeTerminal(sz)
			if err != nil {
				return err
			}
		case message.SessionClose:
			return lc.Close()
		}
	}

	return nil
}
