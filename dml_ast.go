package sqlsmith

import (
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/opcode"
	"github.com/you06/sqlsmith-go/util"
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

func (s *SQLSmith) updateStmt() ast.Node {
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
		updateStmtNode.Where = s.binaryOperationExpr(whereRand)
	} else {
		updateStmtNode.Where = ast.NewValueExpr(1)
	}

	return &updateStmtNode
}

func (s *SQLSmith) insertStmt() ast.Node {
	insertStmtNode := ast.InsertStmt{
		Table: &ast.TableRefsClause{
			TableRefs: &ast.Join{
				Left: &ast.TableName{},
			},
		},
		Lists: [][]ast.ExprNode{},
		Columns: []*ast.ColumnName{},
	}
	return &insertStmtNode
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

func (s *SQLSmith) binaryOperationExpr(depth int) *ast.BinaryOperationExpr {
	node := ast.BinaryOperationExpr{}
	if depth > 0 {
		r := util.Rd(4)
		switch r {
		case 0:
			node.Op = opcode.LogicXor
		case 1:
			node.Op = opcode.LogicOr
		default:
			node.Op = opcode.LogicAnd
		}
		node.L = s.binaryOperationExpr(0)
		node.R = s.binaryOperationExpr(depth - 1)
	} else {
		r := util.Rd(4)
		switch r {
		case 0:
			node.Op = opcode.GT
		case 1:
			node.Op = opcode.LT
		case 2:
			node.Op = opcode.NE
		default:
			node.Op = opcode.EQ
		}
		node.L = &ast.ColumnNameExpr{}
		r = util.Rd(3)
		switch r{
		case 0:
			// hope there is an empty value type
			node.R = ast.NewValueExpr(1)
		default:
			node.R = &ast.ColumnNameExpr{}
		}
	}
	return &node
}
