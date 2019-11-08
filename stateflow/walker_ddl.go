package stateflow

import (
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/model"
	parserTypes "github.com/pingcap/parser/types"
	"github.com/you06/sqlsmith-go/builtin"
	"github.com/you06/sqlsmith-go/types"
	"github.com/you06/sqlsmith-go/util"
)

func (s *StateFlow) walkCreateTableStmt(node *ast.CreateTableStmt) *types.Table {
	table := s.randNewTable()
	for _, column := range table.Columns {
		node.Cols = append(node.Cols, &ast.ColumnDef{
			Name: &ast.ColumnName{
				Name: model.NewCIStr(column.Column),
			},
			Tp: s.makeFieldType(column.DataType, column.DataLen),
			Options: s.makeColumnOptions(column, column.Options),
		})
		if column.HasOption(ast.ColumnOptionAutoIncrement) {
			s.makeConstraintPrimaryKey(node, column)
		}
	}
	node.Table.Name = model.NewCIStr(table.Table)
	s.walkTableOption(node)
	return nil
}

func (s *StateFlow) makeFieldType(t string, l int) *parserTypes.FieldType {
	fieldType := parserTypes.NewFieldType(util.Type2Tp(t))
	fieldType.Flen = l
	return fieldType
}

func (s *StateFlow) makeColumnOptions(column *types.Column, options []ast.ColumnOptionType) (columnOptions []*ast.ColumnOption) {
	for _, opt := range options {
		columnOptions = append(columnOptions, s.makeColumnOption(column, opt))
	}
	return
}

func (s *StateFlow) makeColumnOption(column *types.Column, option ast.ColumnOptionType) *ast.ColumnOption {
	columnOption := ast.ColumnOption{
		Tp: option,
	}
	if option == ast.ColumnOptionDefaultValue {
		columnOption.Expr = builtin.GenerateTypeFuncCallExpr(column.DataType)
	}
	return &columnOption
}

// makeConstraintPromaryKey is for convenience
func (s *StateFlow) makeConstraintPrimaryKey(node *ast.CreateTableStmt, column *types.Column) {
	node.Constraints = append(node.Constraints, &ast.Constraint{
		Tp: ast.ConstraintPrimaryKey,
		Keys: []*ast.IndexColName{
			{
				Column: &ast.ColumnName{
					Name: model.NewCIStr(column.Column),
				},
			},
		},
	})
}

func (s *StateFlow) walkTableOption(node *ast.CreateTableStmt) {
	node.Options = append(node.Options, &ast.TableOption{
		Tp: ast.TableOptionEngine,
		StrValue: "InnoDB",
	})
	node.Options = append(node.Options, &ast.TableOption{
		Tp: ast.TableOptionCharset,
		StrValue: util.RdCharset(),
	})
}
