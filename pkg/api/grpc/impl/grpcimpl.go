package impl

import (
	"encoding/json"
	"io"
	"log"
	"fmt"

	"github.com/btwiuse/invctrl/pkg/api"
	"github.com/btwiuse/wetty/localcmd"
	"github.com/btwiuse/wetty/wetty"
	"github.com/kr/pty"
)

type Slave struct {
	*localcmd.Factory
	Name string
	// client id/index, to distinguish logs of different commands
}

func (bs *Slave) Send(sendServer api.Slave_SendServer) error {
	lc, err := bs.Factory.New()
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
		log.Println(bs.Name, wetty.IotaString(msg.Type[0]), fmt.Sprintf("%q", string(msg.Body)))
		switch msgType := msg.Type[0]; msgType {
		case wetty.Input:
			_, err = lc.Write(msg.Body)
			if err != nil {
				log.Println(bs.Name, "error writing to lc:", err)
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
		case wetty.ClientDead:
			return lc.Close()
		}
	}

	return nil
}
