package sqlsmith

import (
	"github.com/pingcap/parser/ast"
	"math/rand"
	"time"
)

func (s *SQLSmith) selectStmt (depth int) ast.Node {
	selectStmtNode := ast.SelectStmt{}

	if s.rd(10) < 5 {
		selectStmtNode.From = nil
	}

	return &selectStmtNode
}

func (s *SQLSmith) constructSelectStmt(pNode ast.Node, depth int) (cNode ast.Node) {
	if depth<=0{
		return pNode
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100)
	switch pNode.(type) {
	case nil:
		return s.constructSelectStmt(nil, depth - 1)
	case *ast.SelectStmt:
		return s.constructSelectStmt(nil, depth - 1)
	case *ast.TableRefsClause:
		return s.constructSelectStmt(nil, depth - 1)
	case ast.ExprNode:
		return s.constructSelectStmt(nil, depth - 1)
	case *ast.GroupByClause:
		return s.constructSelectStmt(nil, depth - 1)
	case *ast.HavingClause:
		return s.constructSelectStmt(nil, depth - 1)
	case *ast.WindowSpec:
		return s.constructSelectStmt(nil, depth - 1)
	case *ast.OrderByClause:
		return s.constructSelectStmt(nil, depth - 1)
	case *ast.Limit:
		return s.constructSelectStmt(nil, depth - 1)
	case *ast.Join:
		return s.constructSelectStmt(nil, depth - 1)
	case ast.ResultSetNode:
		return s.constructSelectStmt(nil, depth - 1)
	case *ast.OnCondition:
		return s.constructSelectStmt(nil, depth - 1)
	default:
		return pNode
	}
}