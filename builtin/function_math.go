package builtin


import "github.com/pingcap/parser/ast"

var mathFunctions = []*functionClass{
	{ast.Abs, 1, 1, false, true},
	{ast.Acos, 1, 1, false, true},
	{ast.Asin, 1, 1, false, true},
	{ast.Atan, 1, 2, false, true},
	{ast.Atan2, 2, 2, false, true},
	{ast.Ceil, 1, 1, false, true},
	{ast.Ceiling, 1, 1, false, true},
	{ast.Conv, 3, 3, false, true},
	{ast.Cos, 1, 1, false, true},
	{ast.Cot, 1, 1, false, true},
	{ast.CRC32, 1, 1, false, true},
	{ast.Degrees, 1, 1, false, true},
	{ast.Exp, 1, 1, false, true},
	{ast.Floor, 1, 1, false, true},
	{ast.Ln, 1, 1, false, true},
	{ast.Log, 1, 2, false, true},
	{ast.Log2, 1, 1, false, true},
	{ast.Log10, 1, 1, false, true},
	{ast.PI, 0, 0, false, true},
	{ast.Pow, 2, 2, false, true},
	{ast.Power, 2, 2, false, true},
	{ast.Radians, 1, 1, false, true},
	{ast.Rand, 0, 1, false, true},
	{ast.Round, 1, 2, false, true},
	{ast.Sign, 1, 1, false, true},
	{ast.Sin, 1, 1, false, true},
	{ast.Sqrt, 1, 1, false, true},
	{ast.Tan, 1, 1, false, true},
	{ast.Truncate, 2, 2, false, true},
}
