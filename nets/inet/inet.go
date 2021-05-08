package inet

import "net"

type IServer interface {
	StartAt(address string, mgr IAgentManager)
}

type IAgentManager interface {
	OnNewConn(conn net.Conn)
	Get(id uint64) IAgent
	Del(id uint64)
}

type IAgent interface {
	Start()
	GetId() uint64
}

type IAgentHandler interface {
	OnReady()
	// Close()
	OnClose()
	OnRcvData(data []byte)
}
