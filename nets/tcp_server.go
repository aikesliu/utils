package nets

import (
	"net"
	"sync"
	"time"

	"github.com/aikesliu/utils/log"
	"github.com/aikesliu/utils/nets/inet"
)

type TcpServer struct {
	l sync.RWMutex

	address string
	ln      net.Listener

	mgr inet.IAgentManager
}

func (m *TcpServer) StartAt(address string, mgr inet.IAgentManager) {
	if mgr == nil {
		log.E("server IAgentManager is nil")
		return
	}
	m.address = address
	m.mgr = mgr
	ln, err := net.Listen("tcp", m.address)
	log.I("server listen at %v", m.address)
	if err != nil {
		log.F("tcp server listen failed: %v", err)
		return
	}
	m.ln = ln
	go m.run()
}

func (m *TcpServer) run() {
	var tempDelay time.Duration
	for {
		conn, err := m.ln.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.E("accept failed: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return
		}
		tempDelay = 0
		m.mgr.OnNewConn(conn)
	}
}
