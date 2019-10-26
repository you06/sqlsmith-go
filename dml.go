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
		sstmt := &ast.SelectStmt{}
		if r>10{
			sstmt.Distinct = true
		}
		if r>20{
			sstmt.From = &ast.TableRefsClause{
				TableRefs:&ast.Join{

				},
			}
		}
		if r>30{

		}
		if r > 40{

		}
		return s.constructSelectStmt(sstmt, depth - 1)
	case *ast.SelectStmt:
	case *ast.TableRefsClause:
	case ast.ExprNode:
	case *ast.GroupByClause:
	case *ast.HavingClause:
	case *ast.WindowSpec:
	case *ast.OrderByClause:
	case *ast.Limit:
	case *ast.Join:
	case ast.ResultSetNode:
	case *ast.OnCondition:
	default:
		return pNode
	}
}