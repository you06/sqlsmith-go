package stateflow

import (
	"github.com/pingcap/tidb/types/parser_driver"
	tidbTypes "github.com/pingcap/tidb/types"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/model"
	"github.com/you06/sqlsmith-go/types"
	"github.com/you06/sqlsmith-go/builtin"
	"github.com/you06/sqlsmith-go/util"
)

// WalkTree parse
func (s *StateFlow) WalkTree(node ast.Node) ast.Node {
	s.walkTree(node)
	return node
}

func (s *StateFlow) walkTree(node ast.Node) {
	switch node := node.(type) {
	// DML
	case *ast.SelectStmt:
		_ = s.walkSelectStmt(node)
	case *ast.UpdateStmt:
		_ = s.walkUpdateStmt(node)
	case *ast.InsertStmt:
		_ = s.walkInsertStmt(node)
	// DDL
	case *ast.CreateTableStmt:
		_ = s.walkCreateTableStmt(node)
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

		mergeTable, _ := s.mergeTable(lTable, rTable)
		if node.From.TableRefs.On != nil {
			s.walkOnStmt(node.From.TableRefs.On, lTable, rTable)
		}
		table = mergeTable
		s.walkSelectStmtColumns(node, table, true)
	}
	return table
}

func (s *StateFlow) walkUpdateStmt(node *ast.UpdateStmt) *types.Table {
	table := s.walkTableName(node.TableRefs.TableRefs.Left.(*ast.TableName), false)
	for len(table.Columns) == 0 {
		table = s.walkTableName(node.TableRefs.TableRefs.Left.(*ast.TableName), false)
	}
	s.walkAssignmentList(&node.List, table)
	s.walkExprNode(node.Where, table, nil)
	// switch node := node.Where.(type) {
	// case *ast.BinaryOperationExpr:
	// 	s.walkBinaryOperationExpr(node, table)
	// }
	return table
}

func (s *StateFlow) walkInsertStmt(node *ast.InsertStmt) *types.Table {
	table := s.walkTableName(node.Table.TableRefs.Left.(*ast.TableName), false)	
	s.walkColumns(&node.Columns, table)
	s.walkLists(&node.Lists, table)
	return nil
}

func (s *StateFlow) walkOnStmt(node *ast.OnCondition, table1, table2 *types.Table) {
	switch node := node.Expr.(type) {
	case *ast.BinaryOperationExpr:
		s.walkExprNode(node.R, table2, s.walkExprNode(node.L, table1, nil))
	}
	// if node.From.TableRefs.On != nil {
		// if onColumns[0] == nil {
		// 	// node.From.TableRefs.On = ast.
		// }
		// if onColumns[1] == nil {
		// 	// TODO add some builtin function to on clause
		// 	// node.From.TableRefs.On = nil
		// } else {
		// 	switch node := node.Expr.(type) {
		// 	case *ast.BinaryOperationExpr:
		// 		s.walkExprNode(node.L, onColumns[0])
		// 		s.walkExprNode(node.R, onColumns[1])
		// 	}
		// }
	// 		if node, ok := node.From.TableRefs.On.Expr.(*ast.BinaryOperationExpr); ok {
	// 		}
	// 	}
	// }
}

func (s *StateFlow) walkResultSetNode(node ast.ResultSetNode) *types.Table {
	switch node := node.(type) {
	case *ast.TableName:
		return s.walkTableName(node, true)
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

func (s *StateFlow) walkTableName(node *ast.TableName, fn bool) *types.Table {
	table := s.randTable(false, fn)
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
				Expr: builtin.GenerateFuncCallExpr(table, util.Rd(4)),
				AsName: model.NewCIStr(column.Column),
			})
		}
	}
}

func (s *StateFlow) walkExprNode(node ast.ExprNode, table *types.Table, column *types.Column) *types.Column {
	switch node := node.(type) {
	case *ast.BinaryOperationExpr:
		s.walkBinaryOperationExpr(node, table)
	case *ast.ColumnNameExpr:
		return s.walkColumnNameExpr(node, table)
	case *driver.ValueExpr:
		s.walkValueExpr(node, table, column)
	}
	return nil
}

func (s *StateFlow) walkColumnNameExpr(node *ast.ColumnNameExpr, table *types.Table) *types.Column {
	column := table.RandColumn()
	node.Name = &ast.ColumnName{
		Table: model.NewCIStr(column.Table),
		Name: model.NewCIStr(column.Column),
	}
	return column
}

func (s *StateFlow) walkValueExpr(node *driver.ValueExpr, table *types.Table, column *types.Column) *types.Table {
	if column != nil {
		switch column.DataType {
		case "varchar", "text":
			node.SetString(util.GenerateStringItem())
		case "int":
			node.SetInt64(int64(util.GenerateIntItem()))
		case "float":
			node.SetFloat64(util.GenerateFloatItem())
		case "timestamp":
			node.SetMysqlTime(tidbTypes.Time{
				Time: tidbTypes.FromGoTime(util.GenerateDateItem()),
			})
		}
	}
	return nil
}

func (s *StateFlow) walkAssignmentList(list *[]*ast.Assignment, table *types.Table) {
	columns := s.randColumns(table)
	for _, column := range columns {
		// TODO: specify primary key in type Table
		// to avoid this hard coding
		if column.Column == "id" {
			continue
		}
		assignment := ast.Assignment{
			Column: &ast.ColumnName{
				Table: model.NewCIStr(column.Table),
				Name: model.NewCIStr(column.Column),
			},
			Expr: ast.NewValueExpr(util.GenerateDataItem(column.DataType)),
		}
		*list = append(*list, &assignment)
	}
}

func (s *StateFlow) walkBinaryOperationExpr(node *ast.BinaryOperationExpr, table *types.Table) {
	s.walkExprNode(node.R, table, s.walkExprNode(node.L, table, nil))
}

func (s *StateFlow) walkColumns(columns *[]*ast.ColumnName, table *types.Table) {
	for _, column := range table.Columns {
		if column.Column == "id" {
			continue
		}
		*columns = append(*columns, &ast.ColumnName{
			Table: model.NewCIStr(column.Table),
			Name: model.NewCIStr(column.Column),
		})
	}	
}

func (s *StateFlow) walkLists(lists *[][]ast.ExprNode, table *types.Table) {
	count := util.Rd(20)
	for i := 0; i < count; i++ {
		*lists = append(*lists, s.makeList(table))	
	}
}

func (s *StateFlow) makeList(table *types.Table) []ast.ExprNode {
	var list []ast.ExprNode
	for _, column := range table.Columns {
		if column.Column == "id" {
			continue
		}
		list = append(list, ast.NewValueExpr(util.GenerateDataItem(column.DataType)))
	}
	return list
}
