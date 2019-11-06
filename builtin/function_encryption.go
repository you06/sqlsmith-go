package builtin


import "github.com/pingcap/parser/ast"

var encryptionFunctions = []*functionClass{
	{ast.AesDecrypt, 2, 3, false, true},
	{ast.AesEncrypt, 2, 3, false, true},
	{ast.Compress, 1, 1, false, true},
	{ast.Decode, 2, 2, false, true},
	{ast.DesDecrypt, 1, 2, false, true},
	{ast.DesEncrypt, 1, 2, false, true},
	{ast.Encode, 2, 2, false, true},
	{ast.Encrypt, 1, 2, false, true},
	{ast.MD5, 1, 1, false, true},
	{ast.OldPassword, 1, 1, false, true},
	{ast.PasswordFunc, 1, 1, false, true},
	{ast.RandomBytes, 1, 1, false, true},
	{ast.SHA1, 1, 1, false, true},
	{ast.SHA, 1, 1, false, true},
	{ast.SHA2, 2, 2, false, true},
	{ast.Uncompress, 1, 1, false, true},
	{ast.UncompressedLength, 1, 1, false, true},
	{ast.ValidatePasswordStrength, 1, 1, false, true},
}
