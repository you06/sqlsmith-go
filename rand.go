package sqlsmith

func (s *SQLSmith) rd (n int) int {
	return s.Rand.Intn(n)
}

func (s *SQLSmith) rdRange (n, m int) int {
	if m < n {
		n, m = m, n
	}
	return n + s.Rand.Intn(m - n)
}
