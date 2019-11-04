package stateflow

func (s *StateFlow) rd (n int) int {
	return s.rand.Intn(n)
}

func (s *StateFlow) rdRange (n, m int) int {
	if m < n {
		n, m = m, n
	}
	return n + s.rand.Intn(m - n)
}

func (s *StateFlow) rdFloat64() float64 {
	return s.rand.Float64()
}
