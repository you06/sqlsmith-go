package sqlsmith

import (
	"bytes"
	"math/rand"
	"time"

	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/format"
)

func (s *SQLSmith) selectStmt(depth int) ast.Node {
	selectStmtNode := ast.SelectStmt{}

	if s.rd(10) < 5 {
		selectStmtNode.From = nil
	}

	return &selectStmtNode
}

func (s *SQLSmith) constructSelectStmt(pNode ast.Node, depth int) (cNode ast.Node) {
	if depth <= 0 {
		return pNode
	}
	rand.Seed(time.Now().UnixNano())
	_ = rand.Intn(100)
	switch pNode.(type) {
	case nil:
		sstmt := &ast.SelectStmt{}
		s.Node = sstmt
		return s.constructSelectStmt(nil, depth-1)
	case *ast.SelectStmt:
		return s.constructSelectStmt(nil, depth-1)
	case *ast.TableRefsClause:
		return s.constructSelectStmt(nil, depth-1)
	case ast.ExprNode:
		return s.constructSelectStmt(nil, depth-1)
	case *ast.GroupByClause:
		return s.constructSelectStmt(nil, depth-1)
	case *ast.HavingClause:
		return s.constructSelectStmt(nil, depth-1)
	case *ast.WindowSpec:
		return s.constructSelectStmt(nil, depth-1)
	case *ast.OrderByClause:
		return s.constructSelectStmt(nil, depth-1)
	case *ast.Limit:
		return s.constructSelectStmt(nil, depth-1)
	case *ast.Join:
		return s.constructSelectStmt(nil, depth-1)
	case ast.ResultSetNode:
		return s.constructSelectStmt(nil, depth-1)
	case *ast.OnCondition:
		return s.constructSelectStmt(nil, depth-1)
	default:
		return pNode
	}
}

func (s *SQLSmith) ToSQL() (string, error) {
	b := bytes.NewBuffer([]byte{})
	err := s.Node.Restore(
		&format.RestoreCtx{
			Flags:     format.RestoreStringEscapeBackslash,
			In:        b,
			JoinLevel: 10,
		})
	return b.String(), err
}
