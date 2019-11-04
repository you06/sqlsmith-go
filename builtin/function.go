package builtin 

import (
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/model"
	"github.com/you06/sqlsmith-go/types"
)

// GenerateFuncCallExpr generate random builtin chain
func GenerateFuncCallExpr(table *types.Table, args int) *ast.FuncCallExpr {
	funcCallExpr := ast.FuncCallExpr{}

	fns := getValidArgsFunc(args)
	fn := fns[rd(len(fns))]
	funcCallExpr.FnName = model.NewCIStr(fn.name)
	for i := 0; i < args; i++ {
		// r := s.rd(100)
		// if r > 80 {
			
		// }
		funcCallExpr.Args = append(funcCallExpr.Args, GenerateFuncCallExpr(table, 0))
	}

	return &funcCallExpr
}

func getValidArgsFunc(args int) []*functionClass {
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
