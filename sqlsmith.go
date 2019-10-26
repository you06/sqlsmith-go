package sqlsmith

import (
	// "github.com/pingcap/parser"
	"math/rand"
	"time"

	"github.com/pingcap/parser/ast"
)

// New create SQLSmith instance
func New() *SQLSmith {
	return &SQLSmith{
		Rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
		Databases: make(map[string]*Database),
	}
}

// SelectStmt make random select statement
func (s *SQLSmith) SelectStmt(depth int) ast.Node {
	s.depth = depth
	s.maxDepth = 1
	return s.selectStmt(1)
}

// BatchGenSQL generate SQLs in batch
func (s *SQLSmith) BatchGenSQL(n int) []string {
	if n <= 0 {
		return []string{}
	}
	rand.Seed(time.Now().Unix())
	// p := parser.New()
	// p.Parse()
	var (
		result = make([]string, n)
	)
	for i := 0; i < n; i++ {
		_ = s.constructSelectStmt(nil, rand.Intn(10)+5)
		s, err := s.ToSQL()
		if err != nil {
			continue
		}
		result[i] = s
	}
	return result
}
