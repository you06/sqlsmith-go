package sqlsmith

import (
	"github.com/pingcap/parser/ast"
	"math/rand"
)

type SQLSmith struct {
	depth int
	Rand *rand.Rand
	Node ast.Node
}
