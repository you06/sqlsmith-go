package sqlsmith

import (
	"fmt"
	"regexp"
)

const typeRegex = `\(\d+\)`

// LoadSchema init schemas, tables and columns
// record[0] dbname
// record[1] table name
// record[2] table type
// record[3] column name
// record[4] column type
func (s *SQLSmith) LoadSchema (records [][5]string) {
	// init databases
	for _, record := range records {
		dbname := record[0]
		tableName := record[1]
		tableType := record[2]
		columnName := record[3]
		columnType := record[4]
		if _, ok := s.Databases[dbname]; !ok {
			s.Databases[dbname] = &Database{
				Name: dbname,
				Tables: make(map[string]*Table),
			}
		}
		if _, ok := s.Databases[dbname].Tables[tableName]; !ok {
			s.Databases[dbname].Tables[tableName] = &Table{
				DB: dbname,
				Table: tableName,
				Type: tableType,
				Columns: make(map[string]*Column),
			}
		}
		if _, ok := s.Databases[dbname].Tables[tableName].Columns[columnName]; !ok {
			s.Databases[dbname].Tables[tableName].Columns[columnName] = &Column{
				DB: dbname,
				Table: tableName,
				Column: columnName,
				// remove the data size in type definition
				DataType: regexp.MustCompile(typeRegex).ReplaceAllString(columnType, ""),
			}
		}
	}
}

func (s *SQLSmith) randTableFromTable(table *Table, newName bool) (*Table) {
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
		if newName {
			column.Table = newTableName
		} else {
			column.Table = table.Table
		}
		newTable.Columns[column.Column] = column
	}
	return &newTable
}

func (s *SQLSmith) randTable(newName bool) (*Table) {
	tables := s.Databases["community"].Tables
	index := 0
	k := s.rd(len(tables))
	for _, table := range tables {
		if index == k {
			return s.randTableFromTable(table, newName)
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
	table := Table{
		DB: table1.DB,
		Table: s.getSubTableName(),
		Type: "SUB TABLE",
		Columns: make(map[string]*Column),
	}

	var onColumns [2]*Column
	index := 0

	for _, column := range table1.Columns {
		table.Columns[fmt.Sprintf("c%d", index)] = &Column{
			DB: column.DB,
			Table: column.Table,
			DataType: column.DataType,
			Column: fmt.Sprintf("c%d", index),
			OriginColumn: column.Column,
		}
		index++
	}
	for _, column := range table2.Columns {
		table.Columns[fmt.Sprintf("c%d", index)] = &Column{
			DB: column.DB,
			Table: column.Table,
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
