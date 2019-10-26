package sqlsmith

import (
	"bytes"
	"github.com/pingcap/parser/model"

	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/format"
)

func (s *SQLSmith) constructSelectStmt(pNode ast.Node, depth int) ast.Node {
	if depth <= 0 {
		return pNode
	}
	r := s.rd(100)
	switch pNode.(type) {
	case nil:
		sstmt := &ast.SelectStmt{}
		s.Node = sstmt
		if r>10{
			sstmt.Distinct = true
		}
		if r>20{
			sstmt.From = &ast.TableRefsClause{
				TableRefs:&ast.Join{
					Left:  nil, // Field
					Right: nil, // Field
					Tp:ast.CrossJoin,
					On: nil,
					Using: nil, // Field
				},
			}
		}
		if r>30{
			sstmt.OrderBy = &ast.OrderByClause{
				Items: nil, // Field
			}
		}
		if r>40{
			sstmt.Fields = nil // Field
		}
		return s.constructSelectStmt(sstmt, depth-1)
	case *ast.SelectStmt:
		var nextNode ast.Node
		if r>10{
			nextNode = &ast.WindowSpec{
				Name: model.NewCIStr(""),
				Ref: model.NewCIStr(""),
				OrderBy:&ast.OrderByClause{
					Items: nil,
				},
				Frame:&ast.FrameClause{},
			}
		}

		return s.constructSelectStmt(nextNode, depth-1)
	case *ast.TableRefsClause:
		var nextNode ast.Node
		if r>10{
			nextNode = nil
		}

		return s.constructSelectStmt(nextNode, depth-1)
	case *ast.GroupByClause:
		var nextNode ast.Node
		if r>10{
			nextNode = &ast.HavingClause{
				Expr:ast.ExprNode(nil),
			}
		}
		return s.constructSelectStmt(nextNode, depth-1)
	case *ast.WindowSpec:
		var nextNode ast.Node
		if r>10{
			nextNode = ast.ExprNode(nil)
		}
		return s.constructSelectStmt(nextNode, depth-1)
	case *ast.OrderByClause:
		var nextNode ast.Node
		if r>10{
			nextNode = &ast.Limit{
				Count:nil,
				Offset:nil,
			}
		}
		return s.constructSelectStmt(nextNode, depth-1)
	case *ast.Limit:
		var nextNode ast.Node
		if r>10{
			nextNode = nil
		}
		return s.constructSelectStmt(nextNode, depth-1)
	case *ast.Join:
		var nextNode ast.Node
		if r>10{
			_ = nextNode
		}
		if r>20{
			_ = nextNode
		}
		return s.constructSelectStmt(nextNode, depth-1)
	case *ast.OnCondition:
		var nextNode ast.Node
		if r>10{
			_ = nextNode
		}
		if r>20{
			_ = nextNode
		}
		return s.constructSelectStmt(nextNode, depth-1)
	default:
		return pNode
	}
}

// ToSQL translate AST to SQL string
func (s *SQLSmith) ToSQL() (string, error) {
	b := bytes.NewBuffer([]byte{})
	err := s.Node.Restore(
		&format.RestoreCtx{
			Flags:     format.RestoreStringEscapeBackslash,
			In:        b,
			JoinLevel: 10,
		})
	return b.String(), err
}
