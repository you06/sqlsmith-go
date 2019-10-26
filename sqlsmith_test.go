package sqlsmith

import "testing"

func TestSQLSmith_ToSQL(t *testing.T) {
	ss := New()
	ss.LoadSchema([][5]string{
		// tasks table
		[5]string{"community", "tasks", "BASE TABLE", "id", "int(11)"},
		[5]string{"community", "tasks", "BASE TABLE", "season", "int(11)"},
		[5]string{"community", "tasks", "BASE TABLE", "complete_user", "varchar(255)"},
		[5]string{"community", "tasks", "BASE TABLE", "complete_team", "int(11)"},
		[5]string{"community", "tasks", "BASE TABLE", "owner", "varchar(255)"},
		[5]string{"community", "tasks", "BASE TABLE", "repo", "varchar(255)"},
		[5]string{"community", "tasks", "BASE TABLE", "title", "varchar(1023)"},
		[5]string{"community", "tasks", "BASE TABLE", "issue_number", "int(11)"},
		[5]string{"community", "tasks", "BASE TABLE", "pull_number", "int(11)"},
		[5]string{"community", "tasks", "BASE TABLE", "level", "varchar(255)"},
		[5]string{"community", "tasks", "BASE TABLE", "min_score", "int(11)"},
		[5]string{"community", "tasks", "BASE TABLE", "score", "int(11)"},
		[5]string{"community", "tasks", "BASE TABLE", "status", "varchar(255)"},
		[5]string{"community", "picks", "BASE TABLE", "created_at", "timestamp"},
		[5]string{"community", "picks", "BASE TABLE", "expired", "varchar(255)"},
		// picks table
		[5]string{"community", "picks", "BASE TABLE", "id", "int(11)"},
		[5]string{"community", "picks", "BASE TABLE", "season", "int(11)"},
		[5]string{"community", "picks", "BASE TABLE", "task_id", "int(11)"},
		[5]string{"community", "picks", "BASE TABLE", "teamID", "int(11)"},
		[5]string{"community", "picks", "BASE TABLE", "user", "varchar(255)"},
		[5]string{"community", "picks", "BASE TABLE", "pull_number", "int(11)"},
		[5]string{"community", "picks", "BASE TABLE", "status", "varchar(255)"},
		[5]string{"community", "picks", "BASE TABLE", "created_at", "timestamp"},
		[5]string{"community", "picks", "BASE TABLE", "updated_at", "timestamp"},
		[5]string{"community", "picks", "BASE TABLE", "closed_at", "timestamp"},
	})
	// resultSet := ss.BatchGenSQL(1)
	// for _,v:=range resultSet{
	// 	t.Log(v)
	// }
	node := ss.SelectStmt(3)
	t.Log(node)

	// sql, err :=	ss.Walk(node)

	// if err != nil {
	// 	t.Fatalf("walk error %v", err)
	// }
	// t.Log(sql)
}
