package line

type FPoint struct {
	X float64 `json:"x" yaml:"x"`
	Y float64 `json:"y" yaml:"y"`
}

func NewCardinalSpline(points []*FPoint, tension float64) *CardinalSpline {
	if len(points) == 0 {
		return nil
	}
	cs := &CardinalSpline{
		points:  points,
		tension: tension,
	}
	cs.delta = 1 / float64(len(points)-1)
	return cs
}

type CardinalSpline struct {
	points  []*FPoint
	tension float64
	delta   float64
}

func (m *CardinalSpline) GetPoint(idx int) *FPoint {
	if idx < 0 {
		idx = 0
	} else if idx > len(m.points)-1 {
		idx = len(m.points) - 1
	}
	return m.points[idx]
}

func (m *CardinalSpline) GetCardinalSplineAt(idx1, idx2, idx3, idx4 int, t float64) *FPoint {
	t2 := t * t
	t3 := t2 * t
	// log.D("GetCardinalSplineAt, s: %v, t: %v, t2: %v, t3: %v", s, t, t2, t3)

	p1 := m.GetPoint(idx1)
	p2 := m.GetPoint(idx2)
	p3 := m.GetPoint(idx3)
	p4 := m.GetPoint(idx4)

	// log.D("GetCardinalSplineAt, p1: %+v, p2: %+v, p3: %+v, p4: %+v", p1, p2, p3, p4)

	b1 := m.tension * (-t3 + 2*t2 - t)
	b2 := m.tension*(-t3+t2) + (2*t3 - 3*t2 + 1)
	b3 := m.tension*(t3-2*t2+t) + (-2*t3 + 3*t2)
	b4 := m.tension * (t3 - t2)

	x := p1.X*b1 + p2.X*b2 + p3.X*b3 + p4.X*b4
	y := p1.Y*b1 + p2.Y*b2 + p3.Y*b3 + p4.Y*b4

	// log.D("GetCardinalSplineAt, b1: %v, b2: %v, b3: %v, b4: %v", b1, b2, b3, b4)

	return &FPoint{
		X: x,
		Y: y,
	}
}

func (m *CardinalSpline) GetNewPos(percent float64) *FPoint {
	idx := 0
	lt := float64(0)
	if percent < 0 {
		percent = 0
	}
	if percent >= 1 {
		idx = len(m.points) - 1
		lt = 1
	} else {
		fIdx := percent / m.delta
		idx = int(fIdx)
		lt = fIdx - float64(idx)
	}
	pos := m.GetCardinalSplineAt(idx-1, idx, idx+1, idx+2, lt)
	// log.D("dt: %v, p: %v, lt: %v, x: %v, y: %v", percent, idx, lt, pos.X, pos.Y)
	return pos
}
