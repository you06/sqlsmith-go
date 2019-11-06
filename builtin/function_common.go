package builtin


import "github.com/pingcap/parser/ast"

var commonFunctions = []*functionClass{
	{ast.Coalesce, 1, -1, false, true},
	{ast.IsNull, 1, 1, false, true},
	{ast.Greatest, 2, -1, false, true},
	{ast.Least, 2, -1, false, true},
	{ast.Interval, 2, -1, false, true},
}
