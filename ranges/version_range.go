package util

import (
	"fmt"
	"strings"
)

type VersionRange struct {
	// 空表示没有下限
	Min string `yaml:"min" json:"min"`
	// 空表示没有上限
	Max string `yaml:"max" json:"max"`
}

func Compare(ver1, ver2 string) int {
	if ver1 == ver2 {
		return 0
	}
	// 后缀多余的 .0 去掉后再判断
	// 避免 1.0 和 1.0.0 版本不同的结果
	for {
		l1 := len(ver1)
		ver1 = strings.TrimSuffix(ver1, ".0")
		tl1 := len(ver1)

		l2 := len(ver2)
		ver2 = strings.TrimSuffix(ver2, ".0")
		tl2 := len(ver2)

		if l1 == tl1 && l2 == tl2 {
			break
		}
	}
	s1 := strings.Split(ver1, ".")
	s2 := strings.Split(ver2, ".")
	len1 := len(s1)
	len2 := len(s2)

	compareLen := len1
	if compareLen > len2 {
		compareLen = len2
	}
	for i := 0; i < compareLen; i++ {
		ret := strings.Compare(s1[i], s2[i])
		if ret != 0 {
			return ret
		}
	}
	return len1 - len2
}

func (m *VersionRange) String() string {
	return fmt.Sprintf("(%s~%s)", m.Min, m.Max)
}

// 判断版本是否在区间内 [min, max]
//
// ver 为空，表示始终在区间内,
// min 为空表示无下限,
// max 为空表示无上限.
func (m *VersionRange) IsIn(ver string) bool {
	if ver == "" {
		return true
	}
	if m.Min == "" && m.Max == "" {
		return true
	}
	// ver < min ver
	if m.Min != "" && Compare(ver, m.Min) < 0 {
		return false
	}
	// max ver < ver
	if m.Max != "" && Compare(m.Max, ver) < 0 {
		return false
	}
	return true
}

// 判断版本是否在区间内 [min, max)
//
// ver 为空，表示始终在区间内,
// min 为空表示无下限,
// max 为空表示无上限.
func (m *VersionRange) IsInLeft(ver string) bool {
	if m.IsIn(ver) {
		if Compare(m.Max, ver) == 0 {
			return false
		}
		return true
	}
	return false
}
