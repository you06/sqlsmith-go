package sqlsmith

import "fmt"

func (s *SQLSmith) randTableFromTable(table *Table, newName bool, fn bool) (*Table) {
	newTableName := ""
	if newName {
		newTableName = s.getSubTableName()
	}
	columns := s.randColumns(table)
	newTable := Table{
		DB: table.DB,
		Type: table.Type,
		Columns: make(map[string]*Column),
	}
	if newName {
		newTable.OriginTable = table.Table
		newTable.Table = newTableName
	} else {
		newTable.Table = table.Table
	}
	for _, column := range columns {
		newColumn := &Column{
			DB: column.DB,
			Table: column.Table,
			OriginTable: column.OriginTable,
			Column: column.Column,
			OriginColumn: column.OriginColumn,
			DataType: column.DataType,
			Func: column.Func,
		}
		if newName {
			newColumn.Table = newTableName
		} else {
			newColumn.Table = table.Table
		}
		newTable.Columns[newColumn.Column] = newColumn
	}
	if fn {
		r := s.rd(4)
		for i := 0; i < r; i++ {
			k := fmt.Sprintf("f%d", i)
			newTable.Columns[k] = &Column{
				DB: newTable.DB,
				Table: newTable.Table,
				Column: k,
				DataType: "func",
				Func: true,
			}	
		}
	}
	return &newTable
}

func (s *SQLSmith) randTable(newName bool, fn bool) (*Table) {
	tables := s.Databases[s.currDB].Tables
	index := 0
	k := s.rd(len(tables))
	for _, table := range tables {
		if index == k {
			return s.randTableFromTable(table, newName, fn)
		}
		index++
	}
	// should not reach here
	return nil
}

func (s *SQLSmith) randColumns(table *Table) []*Column {
	var columns []*Column
	rate := float64(0.75)
	for _, column := range table.Columns {
		if s.rdFloat64() < rate {
			columns = append(columns, column)
		}
	}
	return columns
}

func (s *SQLSmith) mergeTable(table1 *Table, table2 *Table) (*Table, [2]*Column) {
	subTableName := s.getSubTableName()
	table := Table{
		DB: table1.DB,
		Table: subTableName,
		Type: "SUB TABLE",
		Columns: make(map[string]*Column),
	}

	var onColumns [2]*Column
	index := 0

	for _, column := range table1.Columns {
		table.Columns[fmt.Sprintf("c%d", index)] = &Column{
			DB: column.DB,
			Table: subTableName,
			OriginTable: column.Table,
			DataType: column.DataType,
			Column: fmt.Sprintf("c%d", index),
			OriginColumn: column.Column,
		}
		index++
	}
	for _, column := range table2.Columns {
		table.Columns[fmt.Sprintf("c%d", index)] = &Column{
			DB: column.DB,
			Table: subTableName,
			OriginTable: column.Table,
			DataType: column.DataType,
			Column: fmt.Sprintf("c%d", index),
			OriginColumn: column.Column,
		}
		index++
	}

	for _, column1 := range table1.Columns {
		for _, column2 := range table2.Columns {
			if column1.DataType == column2.DataType {
				onColumns[0] = column1
				onColumns[1] = column2
				return &table, onColumns
			}
		}
	}

	return &table, onColumns
}

func (s *SQLSmith) renameTable(table *Table) *Table {
	newName := s.getSubTableName()
	table.OriginTable = table.Table
	table.Table = newName
	for _, column := range table.Columns {
		column.Table = newName
	}
	return table
}
