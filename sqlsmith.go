package sqlsmith

import (
	"github.com/pingcap/parser/ast"
	"math/rand"
	"time"
)

func New() *SQLSmith {
	return &SQLSmith{
		Rand: rand.New(rand.NewSource(time.Now().UnixNano())),
		Databases: make(map[string]*Database),
	}
}

// SelectStmt make random select statement
func (s *SQLSmith) SelectStmt (depth int) ast.Node {
	s.depth = depth
	return s.selectStmt(1)
}
