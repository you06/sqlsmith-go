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
	r := rand.Intn(100)
	switch pNode.(type) {
	case nil:
		sstmt := &ast.SelectStmt{}
		s.Node = sstmt
		if r>10{
			sstmt.Distinct = true
		}
		if r>20{

		}
		if r>30{

		}
		return s.constructSelectStmt(nil, depth-1)
	case *ast.SelectStmt:
		var nextNode ast.Node
		if r>10{
			_ = nextNode
		}
		if r>20{
			_ = nextNode
		}
		return s.constructSelectStmt(nil, depth-1)
	case *ast.TableRefsClause:
		var nextNode ast.Node
		if r>10{
			_ = nextNode
		}
		if r>20{
			_ = nextNode
		}
		return s.constructSelectStmt(nil, depth-1)
	case ast.ExprNode:
		var nextNode ast.Node
		if r>10{
			_ = nextNode
		}
		if r>20{
			_ = nextNode
		}
		return s.constructSelectStmt(nil, depth-1)
	case *ast.GroupByClause:
		var nextNode ast.Node
		if r>10{
			_ = nextNode
		}
		if r>20{
			_ = nextNode
		}
		return s.constructSelectStmt(nil, depth-1)
	case *ast.HavingClause:
		var nextNode ast.Node
		if r>10{
			_ = nextNode
		}
		if r>20{
			_ = nextNode
		}
		return s.constructSelectStmt(nil, depth-1)
	case *ast.WindowSpec:
		var nextNode ast.Node
		if r>10{
			_ = nextNode
		}
		if r>20{
			_ = nextNode
		}
		return s.constructSelectStmt(nil, depth-1)
	case *ast.OrderByClause:
		var nextNode ast.Node
		if r>10{
			_ = nextNode
		}
		if r>20{
			_ = nextNode
		}
		return s.constructSelectStmt(nil, depth-1)
	case *ast.Limit:
		var nextNode ast.Node
		if r>10{
			_ = nextNode
		}
		if r>20{
			_ = nextNode
		}
		return s.constructSelectStmt(nil, depth-1)
	case *ast.Join:
		var nextNode ast.Node
		if r>10{
			_ = nextNode
		}
		if r>20{
			_ = nextNode
		}
		return s.constructSelectStmt(nil, depth-1)
	case ast.ResultSetNode:
		var nextNode ast.Node
		if r>10{
			_ = nextNode
		}
		if r>20{
			_ = nextNode
		}
		return s.constructSelectStmt(nil, depth-1)
	case *ast.OnCondition:
		var nextNode ast.Node
		if r>10{
			_ = nextNode
		}
		if r>20{
			_ = nextNode
		}
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
