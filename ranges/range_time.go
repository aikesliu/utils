package ranges

import (
	"fmt"
	"time"
)

type RangeTime struct {
	St time.Time `yaml:"st" json:"st"`
	Et time.Time `yaml:"et" json:"et"`
}

func (m *RangeTime) String() string {
	return fmt.Sprintf("(%v~%v)", m.St, m.Et)
}

// [st, et]
func (m *RangeTime) IsIn(t time.Time) bool {
	if t.Sub(m.St) < 0 {
		return false
	}
	if t.Sub(m.Et) > 0 {
		return false
	}
	return true
}
