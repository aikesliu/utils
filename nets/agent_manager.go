package nets

import (
	"net"
	"sync"

	"github.com/aikesliu/utils/log"
	"github.com/aikesliu/utils/nets/inet"
)

type AgentManager struct {
	l sync.RWMutex

	MaxConn uint64
	agents  map[uint64]inet.IAgent

	NewAgent func(conn net.Conn, id uint64) inet.IAgent
}

func (m *AgentManager) Init() {
	m.agents = map[uint64]inet.IAgent{}
	if m.MaxConn == 0 {
		m.MaxConn = 10
	}
}

func (m *AgentManager) lGetConnId() uint64 {
	m.l.RLock()
	defer m.l.RUnlock()
	for id := uint64(1); id < m.MaxConn; id++ {
		a, ok := m.agents[id]
		if !ok || a == nil {
			return id
		}
	}
	log.E("server has conn: %v, max: %v, cannot get conn id", len(m.agents), m.MaxConn)
	return 0
}

func (m *AgentManager) OnNewConn(conn net.Conn) {
	if m.NewAgent == nil {
		log.E("NewAgent is nil")
		return
	}

	id := m.lGetConnId()
	if id == 0 {
		return
	}

	a := m.NewAgent(conn, id)
	if a == nil {
		log.E("NewAgent return nil")
		return
	}

	m.l.Lock()
	m.agents[id] = a
	m.l.Unlock()

	go func() {
		a.Start()
		m.Del(id)
	}()
}

func (m *AgentManager) Get(id uint64) inet.IAgent {
	m.l.RLock()
	defer m.l.RUnlock()
	a, _ := m.agents[id]
	return a
}

func (m *AgentManager) Del(id uint64) {
	m.l.Lock()
	defer m.l.Unlock()

	// a := m.agents[id]
	delete(m.agents, id)
	// return a
}
