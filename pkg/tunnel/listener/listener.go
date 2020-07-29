package listener

import (
	"bufio"
	"io"
	"log"
	"net"
	"time"

	portless "k0s.io/k0s/pkg/tunnel"
)

// 0. make sure hub is running and listening on ws://127.0.0.1:8345
// 1. connect to hub/tunnel and watch for new conns
// 2. create new conn
func Listener(addr, from string) net.Listener {
	// establish conn
	ln := portless.NewTunnel()
	conn0, err := Dial(addr, from)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		defer func() {
			// TODO: panic and try to reconnect
			//       close ln
			// log.Println("srv: hub lost")
			ln.Close()
		}()

		errchan := make(chan error, 2)

		// recv commands
		scanner := bufio.NewScanner(conn0)
		go func() {
			// ensure err is sent when scanner.Scan fails
			defer func() {
				errchan <- nil
			}()
			for scanner.Scan() {
				if scanner.Text() == "CLOSE" {
					return
				}

				conn, err := Dial(addr, from /*unused*/)
				if err != nil {
					errchan <- err
					return
				}
				ln.Push(conn)
			}
		}()

		// keepalive
		ticker := time.Tick(portless.KeepaliveInterval)
		go func() {
			for range ticker {
				_, err := io.WriteString(conn0, "PING\n")
				if err != nil {
					errchan <- err
					return
				}
			}
		}()

		<-errchan
	}()

	return ln
}
