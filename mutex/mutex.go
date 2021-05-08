package mutex

import (
	"sync"
)

type RWMutex struct {
	l sync.RWMutex
}

func (m *RWMutex) RLock() {
	// log.D("[%p] read lock", m)
	m.l.RLock()
}

func (m *RWMutex) RUnlock() {
	// log.D("[%p] read unlock", m)
	m.l.RUnlock()
}

func (m *RWMutex) Lock() {
	// log.D("[%p] lock", m)
	m.l.Lock()
}

func (m *RWMutex) Unlock() {
	// log.D("[%p] unlock", m)
	m.l.Unlock()
}
