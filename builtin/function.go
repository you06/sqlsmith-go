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
	fn := copyFunc(fns[rd(len(fns))])
	funcCallExpr.FnName = model.NewCIStr(fn.name)
	for i := 0; i < args; i++ {
		r := rd(100)
		if r > 80 {
			funcCallExpr.Args = append(funcCallExpr.Args, GenerateFuncCallExpr(table, 1 + rd(3)))	
		}
		funcCallExpr.Args = append(funcCallExpr.Args, GenerateFuncCallExpr(table, 0))
	}
	// if args != 0 {
	// 	log.Println(funcCallExpr)
	// 	for _, arg := range funcCallExpr.Args {
	// 		log.Println(arg)
	// 	}
	// }
	return &funcCallExpr
}

func getValidArgsFunc(args int) []*functionClass {
	var fns []*functionClass
	for _, fn := range getFuncMap() {
		if fn.minArg > args {
			continue
		}
		if fn.maxArg == -1 || fn.maxArg >= args {
			fns = append(fns, fn)
		}
	}
	return fns
}

func copyFunc(fn *functionClass) *functionClass {
	return &functionClass{
		name: fn.name,
		minArg: fn.minArg,
		maxArg: fn.maxArg,
		constArg: fn.constArg,
		mysql: fn.mysql,
	}
}
