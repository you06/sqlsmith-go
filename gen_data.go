package sqlsmith

import (
	"errors"
	"fmt"
	"strings"
)

// BatchData generate testing data by schema in given batch
// return SQLs with insert statement
func (s *SQLSmith) BatchData(total, batchSize int) ([]string, error) {
	if s.currDB == "" {
		return []string{}, errors.New("no selected database")
	}
	database, ok := s.Databases[s.currDB]
	if !ok {
		return []string{}, errors.New("selected database's schema not loaded")
	}

	var sqls []string
	for _, table := range database.Tables {
		var columns []*Column
		var lines [][]string
		count := 0
		for _, column := range table.Columns {
			columns = append(columns, column)
		}
		for i := 0; i < total; i++ {
			var line []string
			for _, column := range columns {
				line = append(line, s.generateDataItem(column.DataType))
			}
			lines = append(lines, line)
			count ++
			if count >= batchSize {
				count = 0
				sqls = append(sqls, makeSQL(table, columns, lines))
				columns = []*Column{}
			}
		}
		if len(lines) != 0 {
			sqls = append(sqls, makeSQL(table, columns, lines))
		}
	}
	return sqls, nil
}

func makeSQL(table *Table, columns []*Column, lines [][]string) string {
	var columnNames []string
	for _, column := range columns {
		if column.Column == "id" {
			continue
		}
		columnNames = append(columnNames, fmt.Sprintf("`%s`", column.Column))
	}
	return fmt.Sprintf("INSERT INTO `%s` (%s) VALUES \n %s",
		table.Table,
		strings.Join(columnNames, ", "),
		strings.Join(mapFn(lines, func(line []string) string {
			return fmt.Sprintf("(%s)", strings.Join(line, ", "))
		}), ",\n"))
}

func mapFn(arr [][]string, fn func([]string) string) []string {
	var res []string
	for _, item := range arr {
		res = append(res, fn(item))
	}
	return res
} 

func (s *SQLSmith) generateDataItem(columnType string) string {
	switch columnType {
	case "varchar":
		return s.generateStringItem()
	case "text":
		return s.generateStringItem()
	case "int":
		return s.generateIntItem()
	case "timestamp":
		return s.generateDateItem()
	}
	return ""
}

func (s *SQLSmith) generateStringItem() string {
	return fmt.Sprintf("\"%s\"", s.rdString(s.rd(100)))
}

func (s *SQLSmith) generateIntItem() string {
	return fmt.Sprintf("%d", s.rd(2147483647))
}

func (s *SQLSmith) generateFloatItem() string {
	return fmt.Sprintf("%f", float64(s.rd(100000)) * s.rdFloat64())
}

func (s *SQLSmith) generateDateItem() string {
	return fmt.Sprintf("\"%s\"", s.rdDate().Format("2006-1-2 15:04:05"))
}
