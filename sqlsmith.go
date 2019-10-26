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

func (s *SQLSmith) BatchGenSQL(n int) []string {
	if n<=0{
		return []string{}
	}
	rand.Seed(time.Now().Unix())
	result := make([]string, n)
	for i:=0;i<n;i++{
		go func(idx int) {
			_ = s.constructSelectStmt(nil, rand.Int())
			s, err := s.ToSQL()
			if err!=nil{
				return
			}
			result[i] = s
		}(i)
	}
	return result
}