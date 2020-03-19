package coverage

type PresentMetric struct {
	Present bool
}

func (m *PresentMetric) Coverage() float32 {
	if m.Present {
		return 1
	}
	return 0
}

func (m *PresentMetric) Reset() {
	m.Present = false
}




