package coverage

type PresentMetric struct {
	Present  bool	`json:"present"`
	Memoized float32 `json:"coverage"`
}

func (m *PresentMetric) Coverage() float32 {
	if m.Present {
		m.Memoized = 1
	} else {
		m.Memoized = 0
	}
	return m.Memoized
}

func (m *PresentMetric) Reset() {
	m.Present = false
}




