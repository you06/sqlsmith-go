package builtin


import "github.com/pingcap/parser/ast"

var controlFunctions = []*functionClass{
	{ast.If, 3, 3, false, true},
	{ast.Ifnull, 3, 3, false, true},
}
