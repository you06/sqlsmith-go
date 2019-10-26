package sqlsmith

import (
	"github.com/pingcap/parser/ast"
	"math/rand"
)

type SQLSmith struct {
	depth int
	Rand *rand.Rand
	Databases map[string]*Database
	Node ast.Node
}

type Column struct {
	DB string
	Table string
	Column string
	DataType string
}

type Table struct {
	DB string
	Table string
	Type string
	Columns map[string]*Column
}

type Database struct {
	Name string
	Tables map[string]*Table
}
