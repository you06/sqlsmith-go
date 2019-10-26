package sqlsmith

import "math/rand"

type SQLSmith struct {
	depth int
	Rand *rand.Rand
	Databases map[string]*Database
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
