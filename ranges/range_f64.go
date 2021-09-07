package util

import "math/rand"

type RangeFloat64 struct {
	Min float64 `yaml:"min" json:"min"`
	Max float64 `yaml:"max" json:"max"`
}

// rand in [min, max)
//
// if min==max, return min
func (m *RangeFloat64) Rand() float64 {
	diff := m.Max - m.Min
	f := rand.Float64()
	return f*diff + m.Min
}

// [min, max]
func (m *RangeFloat64) IsIn(v float64) bool {
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
func (m *RangeFloat64) IsInLeft(v float64) bool {
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
