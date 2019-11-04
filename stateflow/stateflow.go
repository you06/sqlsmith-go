package stateflow

import (
	"math/rand"
	"time"
	"github.com/you06/sqlsmith-go/types"
)

// StateFlow defines the struct
type StateFlow struct {
	// db is a ref took from SQLSmith, should never change it
	db         *types.Database
	rand       *rand.Rand
	tableIndex int
}

// New Create StateFlow
func New(db *types.Database) *StateFlow {
	// copy whole db here may cost time, but ensure the global safety
	// maybe a future TODO
	return &StateFlow{
		db: db,
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
