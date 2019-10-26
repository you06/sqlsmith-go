package sqlsmith

import (
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/model"
	"log"
)

// Walk will walk the tree and fillin tables and columns data
func (s *SQLSmith) Walk(node ast.Node) {
	_ = s.walk(node)
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
			table := s.randTable(false)
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
		
		table, onColumns := s.mergeTable(lTable, rTable)
		s.walkSelectStmtColumns(node, table)
		
		if node.From.TableRefs.On != nil {
			if node, ok := node.From.TableRefs.On.Expr.(*ast.BinaryOperationExpr); ok {
				s.walkExprNode(node.L, onColumns[0])
				s.walkExprNode(node.R, onColumns[1])
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
		table := s.randTable(false)
		log.Println(table)
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
		node.Fields.Fields = append(node.Fields.Fields, &selectField)
	}
}

func (s *SQLSmith) walkExprNode(node ast.ExprNode, column *Column) {
	if node, ok := node.(*ast.ColumnNameExpr); ok {
		node.Name = &ast.ColumnName{
			Table: model.NewCIStr(column.Table),
			Name: model.NewCIStr(column.Column),
		}
	}
}