package sqlsmith

import (
	"github.com/pingcap/parser/ast"
	"github.com/you06/sqlsmith-go/util"
)

func (s *SQLSmith) createTableStmt() ast.Node {
	createTableNode := ast.CreateTableStmt{
		Table: &ast.TableName{},
		Cols: []*ast.ColumnDef{},
		Constraints: []*ast.Constraint{},
		Options: []*ast.TableOption{},
	}
	if util.Rd(4) == 0 {
		s.partitionTable(&createTableNode)
	}
	return &createTableNode
}

func (s *SQLSmith) partitionTable(node *ast.CreateTableStmt) {
	node.Partition = &ast.PartitionOptions{
		PartitionMethod: ast.PartitionMethod{
			ColumnNames: []*ast.ColumnName{},
		},
		Definitions: []*ast.PartitionDefinition{},
	}
}
