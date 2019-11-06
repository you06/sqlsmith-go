package builtin


import "github.com/pingcap/parser/ast"

var controlFunctions = []*functionClass{
	{ast.If, 0, 0, false, true},
	{ast.Ifnull, 0, 0, false, true},
}
