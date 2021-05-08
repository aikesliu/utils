package nets

import (
	"net"

	"github.com/aikesliu/utils/log"
	"github.com/aikesliu/utils/nets/inet"
)

func NewAgentClient(id uint64) *AgentClient {
	a := &AgentClient{
		id: id,
	}
	return a
}

type AgentClient struct {
	id   uint64
	addr string
	*Agent
}

func (m *AgentClient) Dial(addr string, config AgentConfig, h inet.IAgentHandler) error {
	m.addr = addr
	conn, err := net.Dial("tcp", m.addr)
	if err != nil {
		return err
	}
	m.Agent = NewAgent(conn, m.id, h, config)
	return nil
}

func (m *AgentClient) Close() {
	log.D("%v agent client close", m.Agent.Config.Key)
	m.Agent.Close()
}
