package types 

// Column defines database column
type Column struct {
	DB string
	Table string
	OriginTable string
	Column string
	OriginColumn string
	DataType string
	Func bool
	NewFunc bool
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
