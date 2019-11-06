package util

import (
	"bytes"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/format"
)

// BufferOut parser ast node to SQL string
func BufferOut(node ast.Node) (string, error) {
	out := new(bytes.Buffer)
	err := node.Restore(format.NewRestoreCtx(format.RestoreStringDoubleQuotes, out))
	if err != nil {
		return "", err
	}
	return string(out.Bytes()), nil
}
