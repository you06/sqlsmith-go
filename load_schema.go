package sqlsmith

import "regexp"

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
