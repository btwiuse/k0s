package hub

import (
	"fmt"
	"io"
	"log"
	"net/rpc"
	"net/url"
	"strconv"

	"github.com/btwiuse/gods/maps/linkedhashmap"
	"github.com/btwiuse/invctrl/wrap"

	"google.golang.org/grpc"
)

type Agent struct {
	RPCClient      *rpc.Client
	Info           url.Values
	GRPCClientConn chan *grpc.ClientConn
}

type AgentPool struct {
	Agents *linkedhashmap.Map
	Latest *Agent
}

var GlobalAgentPool = &AgentPool{Agents: linkedhashmap.New()}

func (p *AgentPool) Del(uuid string) {
	p.Agents.Remove(uuid)
}

func (p *AgentPool) Get(uuid string) *Agent {
	v, _ := p.Agents.Get(uuid)
	return v.(*Agent)
}

func (p *AgentPool) Add(agent *Agent) {
	p.Agents.Put(agent.Info.Get("id"), agent)
	p.Latest = agent
}

func (p *AgentPool) Dump() {
	log.Println("[agent pool]")
	isCurrent := func(uuid string) string {
		/*
			if (p.Current != nil) && (p.Current.UUID.String() == uuid) {
				return "*"
			}
		*/
		return " "
	}
	for i, v := range p.Agents.Values() {
		slave := v.(*Agent)
		uuid := p.Agents.Keys()[i].(string)
		fmt.Println(
			fmt.Sprintf("[%s]", strconv.Itoa(i+1)),
			isCurrent(uuid),
			uuid,
			slave.Info,
		)
	}
}

func (p *AgentPool) Has(uuid string) bool {
	_, found := p.Agents.Get(uuid)
	return found
}

// we use NewRPCClient over rpc.NewClient(conn)
// so we can remove slave from pool immediately when it is disconnected

/*
          / io.Reader >--->copy()---> io.PipeWriter ===> io.PipeReader = intercepted io.Reader \
net.Conn  - io.Writer \                                                                        wrap.ReadWriteCloser() - rpc.NewClient - *rpc.Client
          \ io.Closer - io.WriteCloser ---------------------------------------------------------
*/
func (agent *Agent) toRPCClient(conn io.ReadWriteCloser) *rpc.Client {
	copy := func(dst io.Writer, src io.Reader) {
		defer agent.onClose()
		if _, err := io.Copy(dst, src); err != nil {
			log.Println(err)
		}
	}
	pr, pw := io.Pipe()
	go copy(pw, conn)
	return rpc.NewClient(wrap.WrapReadWriteCloser(pr, conn))
}

// onclose is called when slave goes offline
// slave.UUID, slave.RemoteAddr, slave.Info
func (agent *Agent) onClose() {
	defer GlobalAgentPool.Dump()
	log.Println("disconnected:", agent.Info.Get("id"))
	GlobalAgentPool.Del(agent.Info.Get("id"))
}
