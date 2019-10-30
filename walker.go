package sqlsmith

import (
	"bytes"
	"log"

	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/model"
	"github.com/pingcap/parser/format"
)

// Walk will walk the tree and fillin tables and columns data
func (s *SQLSmith) Walk(node ast.Node) (string, error) {
	_ = s.walk(node)
	
	out := new(bytes.Buffer)
	err := node.Restore(format.NewRestoreCtx(format.RestoreStringDoubleQuotes, out))
	if err != nil {
		return "", err
	}
	return string(out.Bytes()), nil
}

func (s *SQLSmith) walk(node ast.Node) *Table {
	s.subTableIndex = 0
	switch n := node.(type) {
	case *ast.SelectStmt:
		return s.walkSelectStmt(n)
	}

	return nil
}

func (s *SQLSmith) walkSelectStmt(node *ast.SelectStmt) *Table {
	if node.From.TableRefs.Right == nil && node.From.TableRefs.Left != nil {
		switch node.From.TableRefs.Left.(type) {
		case *ast.TableName:
			// from schema.table
			table := s.randTable(false, true)
			s.walkSelectStmtColumns(node, table)
			if node, ok := node.From.TableRefs.Left.(*ast.TableName); ok {
				node.Name = model.NewCIStr(table.Table)
			}
			return table
		}
	} else if node.From.TableRefs.Right != nil && node.From.TableRefs.Left != nil {
		// from schemaA join schemaB
		lTable := s.walkResultSetNode(node.From.TableRefs.Left)
		
		rTable := s.walkResultSetNode(node.From.TableRefs.Right)
		for rTable.Table == lTable.Table {
			rTable = s.walkResultSetNode(node.From.TableRefs.Right)
		}

		table, onColumns := s.mergeTable(lTable, rTable)
		s.walkSelectStmtColumns(node, table)

		if node.From.TableRefs.On != nil {
			if onColumns[0] == nil || onColumns[1] == nil {
				node.From.TableRefs.On = nil
			} else {
				if node, ok := node.From.TableRefs.On.Expr.(*ast.BinaryOperationExpr); ok {
					s.walkExprNode(node.L, onColumns[0])
					s.walkExprNode(node.R, onColumns[1])
				}
			}
		}
		return table
	}
	// should not walk here
	log.Println("should not walk here")
	log.Println(node)
	return nil
}

func (s *SQLSmith) walkResultSetNode(node ast.ResultSetNode) *Table {
	switch node := node.(type) {
	case *ast.TableName:
		// from schema.table
		table := s.randTable(false, false)
		if table.OriginTable == "" {
			node.Name = model.NewCIStr(table.Table)
		} else {
			node.Name = model.NewCIStr(table.OriginTable)
		}
		return table
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
	// should not walk here
	log.Println("should not walk here")
	log.Println(node)
	return nil
}

func (s *SQLSmith) walkSelectStmtColumns(node *ast.SelectStmt, table *Table) {
	for _, column := range table.Columns {
		// log.Println(column.Table, column.Column)
		if !column.Func {
			var selectField ast.SelectField
			if column.OriginColumn == "" {
				selectField =	ast.SelectField{
					Expr: &ast.ColumnNameExpr{
						Name: &ast.ColumnName{
							Table: model.NewCIStr(column.Table),
							Name: model.NewCIStr(column.Column),
						},
					},
				}
			} else {
				if column.OriginTable != "" {
					selectField =	ast.SelectField{
						AsName: model.NewCIStr(column.Column),
						Expr: &ast.ColumnNameExpr{
							Name: &ast.ColumnName{
								Table: model.NewCIStr(column.OriginTable),
								Name: model.NewCIStr(column.OriginColumn),
							},
						},
					}
				} else {
					selectField =	ast.SelectField{
						AsName: model.NewCIStr(column.Column),
						Expr: &ast.ColumnNameExpr{
							Name: &ast.ColumnName{
								Table: model.NewCIStr(column.Table),
								Name: model.NewCIStr(column.OriginColumn),
							},
						},
					}
				}
			}
			node.Fields.Fields = append(node.Fields.Fields, &selectField)
		} else {
			node.Fields.Fields = append(node.Fields.Fields, &ast.SelectField{
				Expr: s.generateFuncCallExpr(table, s.rd(4)),
				AsName: model.NewCIStr(column.Column),
			})
		}
	}
}

func (s *SQLSmith) walkExprNode(node ast.ExprNode, column *Column) {
	if node, ok := node.(*ast.ColumnNameExpr); ok {
		node.Name = &ast.ColumnName{
			Table: model.NewCIStr(column.Table),
			Name: model.NewCIStr(column.Column),
		}
		// if column.OriginTable != "" {
		// 	node.Name = &ast.ColumnName{
		// 		Table: model.NewCIStr(column.OriginTable),
		// 		Name: model.NewCIStr(column.Column),
		// 	}
		// } else {
		// 	node.Name = &ast.ColumnName{
		// 		Table: model.NewCIStr(column.Table),
		// 		Name: model.NewCIStr(column.Column),
		// 	}
		// }
	}
}
