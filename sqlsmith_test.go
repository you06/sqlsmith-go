package sqlsmith

import "testing"

func TestSQLSmith_ToSQL(t *testing.T) {
	ss := New()
	ss.LoadSchema([][5]string{})
	resultSet := ss.BatchGenSQL(100)
	for _,v:=range resultSet{
		t.Log(v)
	}
}