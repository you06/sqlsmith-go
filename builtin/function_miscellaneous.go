package builtin


import "github.com/pingcap/parser/ast"

var miscellaneousFunctions = []*functionClass{
	{ast.AnyValue, 1, 1, false, true},
	{ast.InetAton, 1, 1, false, true},
	{ast.InetNtoa, 1, 1, false, true},
	{ast.Inet6Aton, 1, 1, false, true},
	{ast.Inet6Ntoa, 1, 1, false, true},
	// {ast.IsFreeLock, 1, 1, false, true},
	{ast.IsIPv4, 1, 1, false, true},
	{ast.IsIPv4Compat, 1, 1, false, true},
	{ast.IsIPv4Mapped, 1, 1, false, true},
	{ast.IsIPv6, 1, 1, false, true},
	{ast.IsUsedLock, 1, 1, false, true},
	// {ast.MasterPosWait, 2, 4, false, true},
	{ast.NameConst, 2, 2, false, true},
	// {ast.ReleaseAllLocks, 0, 0, false, true},
	{ast.UUID, 0, 0, false, true},
	// {ast.UUIDShort, 0, 0, false, true},
}
