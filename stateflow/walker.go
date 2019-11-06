package stateflow

import (
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/model"
	"github.com/you06/sqlsmith-go/types"
	"github.com/you06/sqlsmith-go/builtin"
	"github.com/you06/sqlsmith-go/util"	
)

// WalkTree parse
func (s *StateFlow) WalkTree(node ast.Node) (string, error) {
	s.walkTree(node)
	return util.BufferOut(node)
}

func (s *StateFlow) walkTree(node ast.Node) {
	switch n := node.(type) {
	case *ast.SelectStmt:
		_ = s.walkSelectStmt(n)
	}
}

func (s *StateFlow) walkSelectStmt(node *ast.SelectStmt) *types.Table {
	var table *types.Table
	if node.From.TableRefs.Right == nil && node.From.TableRefs.Left != nil {
		table = s.walkResultSetNode(node.From.TableRefs.Left)
		s.walkSelectStmtColumns(node, table, false)
	} else if node.From.TableRefs.Right != nil && node.From.TableRefs.Left != nil {
		lTable := s.walkResultSetNode(node.From.TableRefs.Left)
		rTable := s.walkResultSetNode(node.From.TableRefs.Right)

		mergeTable, onColumns := s.mergeTable(lTable, rTable)
		if node.From.TableRefs.On != nil {
			s.walkOnStmt(node.From.TableRefs.On, onColumns)
		}
		table = mergeTable
		s.walkSelectStmtColumns(node, table, true)
	}
	return table
}

func (s *StateFlow) walkOnStmt(node *ast.OnCondition, onColumns [2]*types.Column) {
	// if node.From.TableRefs.On != nil {
		if onColumns[0] == nil {
			// node.From.TableRefs.On = ast.
		}
		if onColumns[1] == nil {
			// TODO add some builtin function to on clause
			// node.From.TableRefs.On = nil
		} else {
			switch node := node.Expr.(type) {
			case *ast.BinaryOperationExpr:
				s.walkExprNode(node.L, onColumns[0])
				s.walkExprNode(node.R, onColumns[1])
			}
		}
	// 		if node, ok := node.From.TableRefs.On.Expr.(*ast.BinaryOperationExpr); ok {
	// 		}
	// 	}
	// }
}

func (s *StateFlow) walkResultSetNode(node ast.ResultSetNode) *types.Table {
	switch node := node.(type) {
	case *ast.TableName:
		return s.walkTableName(node)
	case *ast.TableSource:
		n := node
		if node, ok := node.Source.(*ast.SelectStmt); ok {
			table := s.renameTable(s.walkSelectStmt(node))
			if table.OriginTable != "" {
				n.AsName = model.NewCIStr(table.Table)
			}
			return table
		}	
	}
	s.shouldNotWalkHere(node)
	return nil
}

func (s *StateFlow) walkTableName(node *ast.TableName) *types.Table {
	table := s.randTable(false, true)
	node.Schema = model.NewCIStr(table.DB)
	node.Name = model.NewCIStr(table.Table)
	return table
}

func (s *StateFlow) walkSelectStmtColumns(node *ast.SelectStmt, table *types.Table, join bool) {
	for _, column := range table.Columns {
		// log.Println(column.Table, column.Column)
		if !column.Func {
			var selectField ast.SelectField
			if !join && column.OriginColumn == "" {
				selectField =	ast.SelectField{
					Expr: &ast.ColumnNameExpr{
						Name: &ast.ColumnName{
							Table: model.NewCIStr(column.Table),
							Name: model.NewCIStr(column.Column),
						},
					},
				}
			} else {
				selectField =	ast.SelectField{
					AsName: model.NewCIStr(column.Column),
					Expr: &ast.ColumnNameExpr{
						Name: &ast.ColumnName{
							Table: model.NewCIStr(column.OriginTable),
							Name: model.NewCIStr(column.OriginColumn),
						},
					},
				}
			}
			node.Fields.Fields = append(node.Fields.Fields, &selectField)
		} else {
			node.Fields.Fields = append(node.Fields.Fields, &ast.SelectField{
				Expr: builtin.GenerateFuncCallExpr(table, s.rd(4)),
				AsName: model.NewCIStr(column.Column),
			})
		}
	}
}

func (s *StateFlow) walkExprNode(node ast.ExprNode, column *types.Column) {
	switch node := node.(type) {
	case *ast.ColumnNameExpr:
		node.Name = &ast.ColumnName{
			Table: model.NewCIStr(column.Table),
			Name: model.NewCIStr(column.Column),
		}
	}
}
