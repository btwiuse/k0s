package hub

import (
	"io"
	"log"
	"net/rpc"

	"github.com/btwiuse/conntroll/pkg/wrap"
	"github.com/btwiuse/gods/maps/linkedhashmap"

	"google.golang.org/grpc"
)

type Agent struct {
	RPCClient      *rpc.Client           `json:"-"`
	GRPCClientConn chan *grpc.ClientConn `json:"-"`

	// Metadata
	Id        string `json:"id"`
	Connected int64  `json:"connected"`
	Hostname  string `json:"hostname"`
	Whoami    string `json:"whoami"`
	Pwd       string `json:"pwd"`
	OS        string `json:"os"`
	ARCH      string `json:"arch"`
	IP        string `json:"ip"`
	// Info           url.Values
}

type AgentPool struct {
	*linkedhashmap.Map
}

var GlobalAgentPool = &AgentPool{Map: linkedhashmap.New()}

func (p *AgentPool) Del(id string) {
	p.Remove(id)
}

func (p *AgentPool) Get(id string) *Agent {
	v, _ := p.Map.Get(id)
	return v.(*Agent)
}

func (p *AgentPool) Has(id string) bool {
	_, ok := p.Map.Get(id)
	return ok
}

func (p *AgentPool) Add(agent *Agent) {
	p.Put(agent.Id, agent)
}

// we use NewRPCClient over rpc.NewClient(conn)
// so we can remove agent from pool immediately when it is disconnected

/*
                c                           b                  a
          / io.Reader >--->copy()---> io.PipeWriter ===> io.PipeReader = intercepted io.Reader \
net.Conn  - io.Writer \                                                                        wrap.ReadWriteCloser() - rpc.NewClient - *rpc.Client
          \ io.Closer - io.WriteCloser ---------------------------------------------------------
*/
func (agent *Agent) MakeInterceptedRPCClient(c io.ReadWriteCloser) {
	a, b := io.Pipe()
	go func() {
		defer agent.onClose()
		if _, err := io.Copy(b, c); err != nil {
			log.Println(err)
		}
	}()
	agent.RPCClient = rpc.NewClient(wrap.WrapReadWriteCloser(a, c))
}

// onclose is called when agent goes offline
func (agent *Agent) onClose() {
	// TODO: remove Dump
	// panic: runtime error: index out of range [3] with length 3
	// defer GlobalAgentPool.Dump()
	log.Println("disconnected:", agent.Id)
	GlobalAgentPool.Del(agent.Id)
}
