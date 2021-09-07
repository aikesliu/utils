package util

import (
	"fmt"
	"math/rand"
)

type RangeUint32 struct {
	Min uint32 `yaml:"min" json:"min"`
	Max uint32 `yaml:"max" json:"max"`
}

func (m *RangeUint32) String() string {
	return fmt.Sprintf("(%d~%d)", m.Min, m.Max)
}

// rand in [min, max)
func (m *RangeUint32) Rand() uint32 {
	l, u := m.Min, m.Max
	if u < l {
		l, u = u, l
	}
	if u-l == 0 {
		return m.Min
	}
	return rand.Uint32()%(u-l) + l
}

// [min, max]
func (m *RangeUint32) IsIn(v uint32) bool {
	// 只有上限
	if m.Min == 0 && v <= m.Max {
		return true
	}
	// 只有下限
	if m.Min <= v && m.Max == 0 {
		return true
	}
	return m.Min <= v && v <= m.Max
}

// [min, max)
func (m *RangeUint32) IsInLeft(v uint32) bool {
	// 只有上限
	if m.Min == 0 && v < m.Max {
		return true
	}
	// 只有下限
	if m.Min <= v && m.Max == 0 {
		return true
	}
	return m.Min <= v && v < m.Max
}
