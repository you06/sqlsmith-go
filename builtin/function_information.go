package builtin


import "github.com/pingcap/parser/ast"

var informationFunctions = []*functionClass{
	{ast.ConnectionID, 0, 0, false, true},
	{ast.CurrentUser, 0, 0, false, true},
	{ast.CurrentRole, 0, 0, false, true},
	{ast.Database, 0, 0, false, true},
	// This function is a synonym for DATABASE().
	// See http://dev.mysql.com/doc/refman/5.7/en/information-functions.html#function_schema
	{ast.Schema, 0, 0, false, true},
	{ast.FoundRows, 0, 0, false, true},
	{ast.LastInsertId, 0, 1, false, true},
	{ast.User, 0, 0, false, true},
	{ast.Version, 0, 0, false, true},
	{ast.Benchmark, 2, 2, false, true},
	{ast.Charset, 1, 1, false, true},
	{ast.Coercibility, 1, 1, false, true},
	{ast.Collation, 1, 1, false, true},
	{ast.RowCount, 0, 0, false, true},
	{ast.SessionUser, 0, 0, false, true},
	{ast.SystemUser, 0, 0, false, true},
}
