package sqlsmith

import (
	"github.com/pingcap/parser/ast"
)

func (s *SQLSmith) selectStmt (depth int) ast.Node {
	selectStmtNode := ast.SelectStmt{}

	if s.rd(10) < 5 {
		selectStmtNode.From = nil
	}

	return &selectStmtNode
}
