package sqlsmith

import (
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/model"
)

func (s *SQLSmith) generateFuncCallExpr(table *Table, args int) *ast.FuncCallExpr {
	funcCallExpr := ast.FuncCallExpr{}

	fns := s.getValidArgsFunc(args)
	fn := fns[s.rd(len(fns))]
	funcCallExpr.FnName = model.NewCIStr(fn.name)
	for i := 0; i < args; i++ {
		// r := s.rd(100)
		// if r > 80 {
			
		// }
		funcCallExpr.Args = append(funcCallExpr.Args, s.generateFuncCallExpr(table, 0))
	}

	return &funcCallExpr
}

func (s *SQLSmith) getValidArgsFunc(args int) []*functionClass {
	var fns []*functionClass
	for _, fn := range funcs {
		if fn.minArg > args {
			continue
		}
		if fn.maxArg == -1 || fn.maxArg >= args {
			fns = append(fns, fn)
		}
	}
	return fns
}
