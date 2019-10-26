package sqlsmith

import (
	"github.com/pingcap/parser/ast"
	"math/rand"
)

// SQLSmith defines SQLSmith struct
type SQLSmith struct {
	depth int
	maxDepth int
	Rand *rand.Rand
	Databases map[string]*Database
	subTableIndex int
	Node ast.Node
	currDB string
}

// Column defines database column
type Column struct {
	DB string
	Table string
	Column string
	OriginColumn string
	DataType string
}

// Table defines database table
type Table struct {
	DB string
	Table string
	OriginTable string
	Type string
	Columns map[string]*Column
}

// Database defines database database
type Database struct {
	Name string
	Tables map[string]*Table
}
