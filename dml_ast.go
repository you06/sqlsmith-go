package sqlsmith

import (
	// "github.com/pingcap/tidb/types/parser_driver"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/opcode"
)

func (s *SQLSmith) selectStmt(depth int) ast.Node {
	if depth > s.depth {
		s.depth = depth
	}

	selectStmtNode := ast.SelectStmt{
		SelectStmtOpts: &ast.SelectStmtOpts{
			SQLCache: true,
		},
		Fields: &ast.FieldList{
			Fields: []*ast.SelectField{},
		},
	}

	selectStmtNode.From = s.tableRefsClause(depth + 1)

	return &selectStmtNode
}

func (s *SQLSmith) updateStmt(depth int) ast.Node {
	if depth > s.depth {
		s.depth = depth
	}

	updateStmtNode := ast.UpdateStmt{
		List: []*ast.Assignment{},
		TableRefs: &ast.TableRefsClause{
			TableRefs: &ast.Join{
				Left: &ast.TableName{},
			},
		},
	}

	whereRand := s.rd(10)
	if whereRand < 5 {
		updateStmtNode.Where = &ast.BinaryOperationExpr{}
	} else {
		// updateStmtNode.Where = &driver.ValueExpr{}
		updateStmtNode.Where = ast.NewValueExpr(1)
	}

	return &updateStmtNode
}

func (s *SQLSmith) tableRefsClause(depth int) *ast.TableRefsClause {
	if depth > s.depth {
		s.depth = depth
	}

	tableRefsClause := ast.TableRefsClause{
		TableRefs: &ast.Join{
			Left: &ast.TableName{},
		},
	}

	if s.depth < s.maxDepth {
		// if s.rd(100) > 50 {
		// 	tableRefsClause.TableRefs.Right = &ast.TableName{}
		// } else {
		// 	tableRefsClause.TableRefs.Right = &ast.TableSource{
		// 		Source: s.selectStmt(depth + 1),
		// 	}
		// }
		tableRefsClause.TableRefs.Right = &ast.TableSource{
			Source: s.selectStmt(depth + 1),
		}
		if s.rd(100) > 30 {
			tableRefsClause.TableRefs.On = &ast.OnCondition{
				Expr: &ast.BinaryOperationExpr{
					Op: opcode.EQ,
					L: &ast.ColumnNameExpr{},
					R: &ast.ColumnNameExpr{},
				},
			}
		}
	}

	return &tableRefsClause
}
