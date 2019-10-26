package sqlsmith

import (
	"math/rand"
	"sync"
	"time"

	"github.com/pingcap/parser/ast"
)

func New() *SQLSmith {
	return &SQLSmith{
		Rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
		Databases: make(map[string]*Database),
	}
}

// SelectStmt make random select statement
func (s *SQLSmith) SelectStmt(depth int) ast.Node {
	s.depth = depth
	return s.selectStmt(1)
}

func (s *SQLSmith) BatchGenSQL(n int) []string {
	if n <= 0 {
		return []string{}
	}
	rand.Seed(time.Now().Unix())
	var (
		result = make([]string, n)
		wg     = sync.WaitGroup{}
	)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			_ = s.constructSelectStmt(nil, rand.Int())
			s, err := s.ToSQL()
			if err != nil {
				return
			}
			result[i] = s
		}(i)
	}
	wg.Wait()
	return result
}
