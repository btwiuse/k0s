package dial

import (
	"fmt"
	"net"
	"net/url"

	"github.com/btwiuse/conntroll/pkg/agent/config"
)

func WithInfo(hostport string, info url.Values) (net.Conn, error) {
	requestText := fmt.Sprintf("GET %s?%s HTTP/1.0\r\n\r\n", config.Default.DialBasePath, info.Encode())

	conn, err := net.Dial("tcp", hostport)
	if err != nil {
		return nil, err
	}

	_, err = conn.Write([]byte(requestText))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
