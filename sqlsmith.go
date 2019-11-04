package sqlsmith

import (
	"math/rand"
	"time"

	"github.com/pingcap/parser/ast"
	"github.com/you06/sqlsmith-go/stateflow"
	"github.com/you06/sqlsmith-go/types"
)

// SQLSmith defines SQLSmith struct
type SQLSmith struct {
	depth int
	maxDepth int
	Rand *rand.Rand
	Databases map[string]*types.Database
	subTableIndex int
	Node ast.Node
	currDB string
}

// New create SQLSmith instance
func New() *SQLSmith {
	return &SQLSmith{
		Rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
		Databases: make(map[string]*types.Database),
	}
}

// SetDB set currerent database
func (s *SQLSmith) SetDB(db string) {
	s.currDB = db
}

// SelectStmt make random select statement SQL
func (s *SQLSmith) SelectStmt(depth int) (string, error) {
	s.depth = 1
	s.maxDepth = depth
	tree := s.selectStmt(1)
	return stateflow.New(s.Databases[s.currDB]).WalkTree(tree)
}

// Walk will walk the tree and fillin tables and columns data
func (s *SQLSmith) Walk(tree ast.Node) (string, error) {
	return stateflow.New(s.Databases[s.currDB]).WalkTree(tree)
}