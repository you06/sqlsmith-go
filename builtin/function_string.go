package builtin


import "github.com/pingcap/parser/ast"

var stringFunctions = []*functionClass{
	{ast.ASCII, 1, 1, false, true},
	{ast.Bin, 1, 1, false, true},
	{ast.Concat, 1, -1, false, true},
	{ast.ConcatWS, 2, -1, false, true},
	{ast.Convert, 2, 2, false, true},
	{ast.Elt, 2, -1, false, true},
	{ast.ExportSet, 3, 5, false, true},
	{ast.Field, 2, -1, false, true},
	{ast.Format, 2, 3, false, true},
	{ast.FromBase64, 1, 1, false, true},
	{ast.InsertFunc, 4, 4, false, true},
	{ast.Instr, 2, 2, false, true},
	{ast.Lcase, 1, 1, false, true},
	{ast.Left, 2, 2, false, true},
	{ast.Right, 2, 2, false, true},
	{ast.Length, 1, 1, false, true},
	{ast.LoadFile, 1, 1, false, true},
	{ast.Locate, 2, 3, false, true},
	{ast.Lower, 1, 1, false, true},
	{ast.Lpad, 3, 3, false, true},
	{ast.LTrim, 1, 1, false, true},
	{ast.Mid, 3, 3, false, true},
	{ast.MakeSet, 2, -1, false, true},
	{ast.Oct, 1, 1, false, true},
	{ast.OctetLength, 1, 1, false, true},
	{ast.Ord, 1, 1, false, true},
	{ast.Position, 2, 2, false, true},
	{ast.Quote, 1, 1, false, true},
	{ast.Repeat, 2, 2, false, true},
	{ast.Replace, 3, 3, false, true},
	{ast.Reverse, 1, 1, false, true},
	{ast.RTrim, 1, 1, false, true},
	{ast.Space, 1, 1, false, true},
	{ast.Strcmp, 2, 2, false, true},
	{ast.Substring, 2, 3, false, true},
	{ast.Substr, 2, 3, false, true},
	{ast.SubstringIndex, 3, 3, false, true},
	{ast.ToBase64, 1, 1, false, true},
	// not support fully trim, since the trim expression is special
	// SELECT TRIM('miyuri' FROM 'miyuri-jinja-no-miyuri');
	// {ast.Trim, 1, 3, false, true},
	// {ast.Trim, 1, 1, false, true},	
	{ast.Upper, 1, 1, false, true},
	{ast.Ucase, 1, 1, false, true},
	{ast.Hex, 1, 1, false, true},
	{ast.Unhex, 1, 1, false, true},
	{ast.Rpad, 3, 3, false, true},
	{ast.BitLength, 1, 1, false, true},
	{ast.CharFunc, 2, -1, false, true},
	{ast.CharLength, 1, 1, false, true},
	{ast.CharacterLength, 1, 1, false, true},
	{ast.FindInSet, 2, 2, false, true},
}