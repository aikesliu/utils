package util

import (
	"fmt"
	"math/rand"
	"time"
)

type RangeDuration struct {
	Min time.Duration `yaml:"min" json:"min"`
	Max time.Duration `yaml:"max" json:"max"`
}

func (m *RangeDuration) String() string {
	return fmt.Sprintf("(%v~%v)", m.Min, m.Max)
}

func (m *RangeDuration) Rand() time.Duration {
	l, u := m.Min, m.Max
	if u < l {
		l, u = u, l
	}
	if u-l == 0 {
		return m.Min
	}
	t := rand.Int63()%int64(u-l) + int64(l)
	return time.Duration(t)
}

// [min, max)
func (m *RangeDuration) IsInLeft(t time.Duration) bool {
	if m.Min == 0 && m.Max == 0 {
		return true
	}
	if m.Max == 0 && m.Min <= t {
		return true
	}
	if m.Min == 0 && t < m.Max {
		return true
	}
	return m.Min <= t && t < m.Max
}
