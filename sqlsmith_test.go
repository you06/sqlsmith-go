package sqlsmith

import "testing"

func TestSQLSmith_ToSQL(t *testing.T) {
	ss := New()
	ss.LoadSchema([][5]string{
		[5]string{"test", "balances", "BASE TABLE", "id", "int(11)"},
		[5]string{"test", "balances", "BASE TABLE", "user", "varchar(255)"},
		[5]string{"test", "balances", "BASE TABLE", "money", "int(11)"},
		[5]string{"test", "records", "BASE TABLE", "id", "int(11)"},
		[5]string{"test", "records", "BASE TABLE", "from_id", "int(11)"},
		[5]string{"test", "records", "BASE TABLE", "to_id", "int(11)"},
		[5]string{"test", "records", "BASE TABLE", "created_at", "timestamp"},
	})
	resultSet := ss.BatchGenSQL(100)
	for _,v:=range resultSet{
		t.Log(v)
	}
}