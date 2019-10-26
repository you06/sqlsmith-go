package sqlsmith

import "fmt"

func (s *SQLSmith) rd (n int) int {
	return s.Rand.Intn(n)
}

func (s *SQLSmith) rdRange (n, m int) int {
	if m < n {
		n, m = m, n
	}
	return n + s.Rand.Intn(m - n)
}

func (s *SQLSmith) rdFloat64() float64 {
	return s.Rand.Float64()
}

func (s *SQLSmith) getSubTableName() string {
	name := fmt.Sprintf("ss_sub_%d", s.subTableIndex)
	s.subTableIndex++
	return name
}
