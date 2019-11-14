package types

import (
	"github.com/you06/sqlsmith-go/util"
)

// Table defines database table
type Table struct {
	DB string
	Table string
	OriginTable string
	Type string
	Columns map[string]*Column
}

// RandColumn rand column from table
func (t *Table) RandColumn() *Column {
	if len(t.Columns) == 0 {
		return nil
	}
	rdIndex := util.RdRange(0, len(t.Columns))
	index := 0
	for _, column := range t.Columns {
		if rdIndex == index {
			return column.Clone()
		}
		index++
	}
	// should not reach here
	return nil
}
