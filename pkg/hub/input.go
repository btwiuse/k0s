// +build input

package hub

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/readline.v1"

	rpcimpl "github.com/btwiuse/invctrl/pkg/api/rpc/impl"
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

			bash := func(line string, client *Slave) {
				req := rpcimpl.Request{
					Command: line,
				}
				resp := new(rpcimpl.Response)
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
						v := GlobalSlavePool.Slaves.Values()[0]
						client := v.(*Slave)
					*/
					// client := GlobalSlavePool.GetRandom()
					// go echo(strconv.Itoa(i), client)
				}
				/*
					for _, v := range SlavePool.Slaves.Values() {
						client := v.(*Slave)
						go echo(line, client)
					}
				*/
				continue
			case strings.HasPrefix(line, "!map "):
				line = strings.TrimPrefix(line, "!map ")
				for _, v := range GlobalSlavePool.Slaves.Values() {
					client := v.(*Slave)
					go bash(line, client)
				}
				continue
			case line == "":
				continue
			case line == "Exit", line == "Quit":
				os.Exit(0)
			case line == "Ls":
				GlobalSlavePool.Dump()
				continue
			case GlobalSlavePool.Has(line):
				GlobalSlavePool.Current = GlobalSlavePool.Get(line)
				log.Println("current client:", GlobalSlavePool.Current.UUID)
				continue
			default:
				if GlobalSlavePool.Current == nil {
					fmt.Println("[INFO] Your current client is empty. Enter the uuid to the client you want to talk to first:")
					continue
				}
			}

			/*
				if line == "N" {
					client := GlobalSlavePool.Current
					conn, err := client.Pool.Get()
					log.Println("[POOL Size]", client.Pool.Len())
					if err == nil {
						go io.Copy(os.Stderr, conn)
					} else {
						log.Println(err)
					}
				} else {
					go bash(line, GlobalSlavePool.Current)
				}
			*/

			promptNum += 1
		}

		log.Println("stdin input closed")
	}
}
