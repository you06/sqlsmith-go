package builtin


import "github.com/pingcap/parser/ast"

var jsonFunctions = []*functionClass{
	{ast.JSONType, 1, 1, false, true},
	{ast.JSONExtract, 2, -1, false, true},
	{ast.JSONUnquote, 1, 1, false, true},
	{ast.JSONSet, 3, -1, false, true},
	{ast.JSONInsert, 3, -1, false, true},
	{ast.JSONReplace, 3, -1, false, true},
	{ast.JSONRemove, 2, -1, false, true},
	{ast.JSONMerge, 2, -1, false, true},
	{ast.JSONObject, 0, -1, false, true},
	{ast.JSONArray, 0, -1, false, true},
	{ast.JSONContains, 2, 3, false, true},
	{ast.JSONContainsPath, 3, -1, false, true},
	{ast.JSONValid, 1, 1, false, true},
	{ast.JSONArrayAppend, 3, -1, false, true},
	{ast.JSONArrayInsert, 3, -1, false, true},
	{ast.JSONMergePatch, 2, -1, false, true},
	{ast.JSONMergePreserve, 2, -1, false, true},
	{ast.JSONPretty, 1, 1, false, true},
	{ast.JSONQuote, 1, 1, false, true},
	{ast.JSONSearch, 3, -1, false, true},
	{ast.JSONStorageSize, 1, 1, false, true},
	{ast.JSONDepth, 1, 1, false, true},
	{ast.JSONKeys, 1, 2, false, true},
	{ast.JSONLength, 1, 2, false, true},
}
