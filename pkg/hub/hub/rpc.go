package hub

import (
	"fmt"
	"io"
	"net"
	"time"

	types "github.com/btwiuse/conntroll/pkg/hub"
	"github.com/btwiuse/conntroll/pkg/rng"
	"github.com/btwiuse/conntroll/pkg/uuid"
)

var (
	_ types.RPC = (*YS)(nil)
)

func ToRPC(conn net.Conn) types.RPC {
	return &YS{
		id:      uuid.New(),
		name:    rng.New(),
		created: time.Now(),
		Conn:    conn,
	}
}

func (ys *YS) Time() time.Time {
	return ys.created
}

func (ys *YS) Name() string {
	return ys.name
}

func (ys *YS) ID() string {
	return ys.id
}

func (ys *YS) RemoteIP() string {
	ip, _, _ := net.SplitHostPort(ys.Conn.RemoteAddr().String())
	return ip
}

type YS struct {
	id      string
	name    string
	created time.Time
	net.Conn
}

func (ys *YS) NewSession() {
	io.WriteString(ys.Conn, fmt.Sprintln("ACCEPT"))
}
