package util

import (
	"math/rand"
)

type RangeInt64 struct {
	Min int64 `yaml:"min" json:"min"`
	Max int64 `yaml:"max" json:"max"`
}

// rand in [min, max)
func (m *RangeInt64) Rand() int64 {
	l, u := m.Min, m.Max
	if u < l {
		l, u = u, l
	}
	if u-l == 0 {
		return m.Min
	}
	t := rand.Int63()%(u-l) + l
	return t
}

// [min, max]
func (m *RangeInt64) IsIn(v int64) bool {
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
func (m *RangeInt64) IsInLeft(v int64) bool {
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
