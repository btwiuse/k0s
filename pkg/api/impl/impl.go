package impl

import (
	"encoding/json"
	"io"
	"log"

	"github.com/btwiuse/invctrl/pkg/api"
	"github.com/btwiuse/wetty/localcmd"
	"github.com/btwiuse/wetty/wetty"
	"github.com/kr/pty"
)

type BidiStream struct{}

func (bs *BidiStream) Send(sendServer api.BidiStream_SendServer) error {
	lc, err := localcmd.NewLc([]string{"htop"})
	if err != nil {
		return err
	}

	// send
	go func() {
		buf := make([]byte, 1<<16)
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
			outputMsg := &api.Message{Type: []byte{wetty.Output}, Body: buf[:n]}
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
		// log.Println(msg.Type, msg.Body)
		log.Println(msg.Type, len(msg.Body))
		switch msgType := msg.Type[0]; msgType {
		case wetty.Input:
			_, err = lc.Write(msg.Body)
			if err != nil {
				log.Println("error writing to lc:", err)
				return err
			}
		case wetty.ResizeTerminal:
			sz := &pty.Winsize{}
			err = json.Unmarshal(msg.Body, sz)
			if err != nil {
				return err
			}
			err = lc.ResizeTerminal(sz)
			if err != nil {
				return err
			}
			log.Println("new sz:", sz)
		}
	}

	return nil
}
