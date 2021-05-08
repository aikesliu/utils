package nets

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/aikesliu/utils/log"
	"github.com/aikesliu/utils/mutex"
	"github.com/aikesliu/utils/nets/inet"
)

var (
	DefaultAgentConfig = AgentConfig{
		HeadLen:      2,
		Order:        binary.LittleEndian,
		SendCacheLen: 100,
	}
)

type AgentConfig struct {
	Key          string
	HeadLen      int
	Order        binary.ByteOrder
	SendCacheLen int

	RecvTimeOut time.Duration
}

func NewAgent(conn net.Conn, id uint64, h inet.IAgentHandler, config AgentConfig) *Agent {
	a := &Agent{
		id:     id,
		h:      h,
		conn:   conn,
		Config: config,
	}
	// if a.Config.RecvTimeOut == 0 {
	// 	a.Config.RecvTimeOut = time.Minute
	// }
	if a.Config.HeadLen == 0 {
		a.Config.HeadLen = DefaultAgentConfig.HeadLen
	}
	if a.Config.Order == nil {
		a.Config.Order = DefaultAgentConfig.Order
	}
	if a.Config.SendCacheLen == 0 {
		a.Config.SendCacheLen = DefaultAgentConfig.SendCacheLen
	}
	a.Config.Key = fmt.Sprintf("|ci:%v|", a.id)
	return a
}

type Agent struct {
	id   uint64
	conn net.Conn

	l  mutex.RWMutex
	wg sync.WaitGroup
	h  inet.IAgentHandler

	Config AgentConfig

	chSendDataCache chan []byte
	chSendStop      chan bool
}

func (m *Agent) Start() {
	if m.h == nil {
		log.E("%v agent conn IAgentHandler should not be nil", m.Config.Key)
		return
	}
	m.chSendDataCache = make(chan []byte, m.Config.SendCacheLen)
	m.chSendStop = make(chan bool, 1)

	m.wg.Add(2)
	go m.send()
	go m.rcv()
	m.h.OnReady()
	m.wg.Wait()
	m.Close()
}

func (m *Agent) GetId() uint64 {
	return m.id
}

func (m *Agent) Write2Cache(data []byte) {
	m.chSendDataCache <- data
}

func (m *Agent) send() {
	defer func() {
		m.wg.Done()
		log.I("%v send loop stop", m.Config.Key)
	}()
	log.I("%v send loop start", m.Config.Key)
	for {
		select {
		case data, ok := <-m.chSendDataCache:
			if !ok {
				log.E("%v write chan read msg data failed: chan closed", m.Config.Key)
				return
			}
			_, err := m.conn.Write(data)
			if err != nil {
				log.E("%v send failed: %v", m.Config.Key, err)
				return
			}
		case <-m.chSendStop:
			return
		}
	}
}

func (m *Agent) rcv() {
	defer func() {
		m.wg.Done()
		log.I("%v rcv loop stop", m.Config.Key)
	}()
	log.I("%v rcv loop start", m.Config.Key)
	for {
		if m.Config.RecvTimeOut > 0 {
			_ = m.conn.SetDeadline(time.Now().Add(m.Config.RecvTimeOut))
		}
		data, err := m.read()
		if err != nil {
			log.I("%v read message failed: %v", m.Config.Key, err)
			m.chSendStop <- true
			break
		}
		if len(data) == 0 {
			continue
		}
		m.h.OnRcvData(data)
	}
}

func (m *Agent) read() ([]byte, error) {
	bufMsgLen := make([]byte, m.Config.HeadLen)
	if _, err := io.ReadFull(m.conn, bufMsgLen); err != nil {
		return nil, err
	}
	var msgLen uint32
	switch m.Config.HeadLen {
	case 1:
		msgLen = uint32(bufMsgLen[0])
	case 2:
		msgLen = uint32(m.Config.Order.Uint16(bufMsgLen))
	case 4:
		msgLen = m.Config.Order.Uint32(bufMsgLen)
	}
	// check len 过长或过段的协议都不返回 error, 仅打印日志, 丢弃异常消息
	// if !m.isMsgLenValid(msgLen) {
	// 	log.Debug("msg len: %v invalid", msgLen)
	// 	return nil, nil
	// }
	// log.D("%v head len: %v, msg len: %v", m.Config.Key, m.Config.HeadLen, msgLen)
	data := make([]byte, msgLen)
	if _, err := io.ReadFull(m.conn, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (m *Agent) Close() {
	log.I("%v conn close", m.Config.Key)
	m.l.Lock()
	defer m.l.Unlock()
	if m.chSendDataCache != nil {
		close(m.chSendDataCache)
		m.chSendDataCache = nil
	}
	if m.chSendStop != nil {
		close(m.chSendStop)
		m.chSendStop = nil
	}
	m.h.OnClose()
}
