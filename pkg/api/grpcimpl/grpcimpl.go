package grpcimpl

import (
	"encoding/json"
	"io"
	"log"

	"github.com/btwiuse/wetty/pkg/msg"
	"k0s.io/conntroll/pkg/agent"
	"k0s.io/conntroll/pkg/api"

	"go.uber.org/zap"
)

type Session struct {
	ReadOnly   bool
	TtyFactory agent.TtyFactory
	// client id/index, to distinguish logs of different commands
}

func (session *Session) Send(sendServer api.Session_SendServer) error {
	recorder, _ := zap.NewProduction()
	defer recorder.Sync()
	tty, err := session.TtyFactory.MakeTty()
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
			n, err := tty.Read(buf)
			if err == io.EOF {
				return // nil
			}
			if err != nil {
				return // err
			}
			var (
				msgType = msg.Type_SESSION_OUTPUT
				msgBody = buf[:n]
				req     = &api.Message{
					Type: msgType,
					Body: msgBody,
				}
			)
			err = sendServer.Send(req)
			if err != nil {
				return // err
			}
			// log.Println(req.Type, fmt.Sprintf("%q", string(req.Body)))
			/* this causes infinite log loop, be careful
			recorder.Info("send",
				zap.String("type", req.Type.String()),
				zap.String("content", string(req.Body)),
			)
			*/
		}
		return // nil
	}()

	// recv
	for {
		resp, err := sendServer.Recv()
		if err != nil {
			return nil
		}
		// log.Println(resp.Type, fmt.Sprintf("%q", string(resp.Body)))
		recorder.Info("recv",
			zap.String("type", resp.Type.String()),
			zap.String("content", string(resp.Body)),
		)
		switch resp.Type {
		case msg.Type_CLIENT_INPUT:
			if session.ReadOnly {
				break
			}
			_, err = tty.Write(resp.Body)
			if err != nil {
				log.Println("error writing to tty:", err)
				return err
			}
		case msg.Type_SESSION_RESIZE:
			type Winsize struct {
				Rows int
				Cols int
			}
			sz := &Winsize{}
			err = json.Unmarshal(resp.Body, sz)
			if err != nil {
				return err
			}
			err = tty.Resize(sz.Rows, sz.Cols)
			if err != nil {
				return err
			}
		case msg.Type_SESSION_CLOSE:
			return tty.Close()
		}
	}

	return nil
}
