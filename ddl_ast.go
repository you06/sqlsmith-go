package sqlsmith

import (
	"github.com/pingcap/parser/ast"
)

func (s *SQLSmith) createTableStmt() ast.Node {
	createTableNode := ast.CreateTableStmt{
		Table: &ast.TableName{},
		Cols: []*ast.ColumnDef{},
		Constraints: []*ast.Constraint{},
		Options: []*ast.TableOption{},
	}

	return &createTableNode
}
