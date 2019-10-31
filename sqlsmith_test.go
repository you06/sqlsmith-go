package sqlsmith

import "testing"

func TestSQLSmith_ToSQL(t *testing.T) {
	ss := New()
	schema := [][5]string{
		{"community", "comments", "BASE TABLE", "id", "int(11)"},
		{"community", "comments", "BASE TABLE", "owner", "varchar(255)"},
		{"community", "comments", "BASE TABLE", "repo", "varchar(255)"},
		{"community", "comments", "BASE TABLE", "comment_id", "int(11)"},
		{"community", "comments", "BASE TABLE", "comment_type", "varchar(128)"},
		{"community", "comments", "BASE TABLE", "pull_number", "int(11)"},
		{"community", "comments", "BASE TABLE", "body", "text"},
		{"community", "comments", "BASE TABLE", "user", "varchar(255)"},
		{"community", "comments", "BASE TABLE", "url", "varchar(1023)"},
		{"community", "comments", "BASE TABLE", "association", "varchar(255)"},
		{"community", "comments", "BASE TABLE", "relation", "varchar(255)"},
		{"community", "comments", "BASE TABLE", "created_at", "timestamp"},
		{"community", "comments", "BASE TABLE", "updated_at", "timestamp"},
		{"community", "picks", "BASE TABLE", "id", "int(11)"},
		{"community", "picks", "BASE TABLE", "season", "int(11)"},
		{"community", "picks", "BASE TABLE", "task_id", "int(11)"},
		{"community", "picks", "BASE TABLE", "teamID", "int(11)"},
		{"community", "picks", "BASE TABLE", "user", "varchar(255)"},
		{"community", "picks", "BASE TABLE", "pull_number", "int(11)"},
		{"community", "picks", "BASE TABLE", "status", "varchar(128)"},
		{"community", "picks", "BASE TABLE", "created_at", "timestamp"},
		{"community", "picks", "BASE TABLE", "updated_at", "timestamp"},
		{"community", "picks", "BASE TABLE", "closed_at", "timestamp"},
		{"community", "pulls", "BASE TABLE", "id", "int(11)"},
		{"community", "pulls", "BASE TABLE", "owner", "varchar(255)"},
		{"community", "pulls", "BASE TABLE", "repo", "varchar(255)"},
		{"community", "pulls", "BASE TABLE", "pull_number", "int(11)"},
		{"community", "pulls", "BASE TABLE", "title", "text"},
		{"community", "pulls", "BASE TABLE", "body", "text"},
		{"community", "pulls", "BASE TABLE", "user", "varchar(255)"},
		{"community", "pulls", "BASE TABLE", "association", "varchar(255)"},
		{"community", "pulls", "BASE TABLE", "relation", "varchar(255)"},
		{"community", "pulls", "BASE TABLE", "label", "text"},
		{"community", "pulls", "BASE TABLE", "status", "varchar(128)"},
		{"community", "pulls", "BASE TABLE", "created_at", "timestamp"},
		{"community", "pulls", "BASE TABLE", "updated_at", "timestamp"},
		{"community", "pulls", "BASE TABLE", "closed_at", "timestamp"},
		{"community", "pulls", "BASE TABLE", "merged_at", "timestamp"},
		{"community", "tasks", "BASE TABLE", "id", "int(11)"},
		{"community", "tasks", "BASE TABLE", "season", "int(11)"},
		{"community", "tasks", "BASE TABLE", "complete_user", "varchar(255)"},
		{"community", "tasks", "BASE TABLE", "complete_team", "int(11)"},
		{"community", "tasks", "BASE TABLE", "owner", "varchar(255)"},
		{"community", "tasks", "BASE TABLE", "repo", "varchar(255)"},
		{"community", "tasks", "BASE TABLE", "title", "varchar(2047)"},
		{"community", "tasks", "BASE TABLE", "issue_number", "int(11)"},
		{"community", "tasks", "BASE TABLE", "pull_number", "int(11)"},
		{"community", "tasks", "BASE TABLE", "level", "varchar(255)"},
		{"community", "tasks", "BASE TABLE", "min_score", "int(11)"},
		{"community", "tasks", "BASE TABLE", "score", "int(11)"},
		{"community", "tasks", "BASE TABLE", "status", "varchar(255)"},
		{"community", "tasks", "BASE TABLE", "created_at", "timestamp"},
		{"community", "tasks", "BASE TABLE", "expired", "varchar(255)"},
		{"community", "teams", "BASE TABLE", "id", "int(11)"},
		{"community", "teams", "BASE TABLE", "season", "int(11)"},
		{"community", "teams", "BASE TABLE", "name", "varchar(255)"},
		{"community", "teams", "BASE TABLE", "issue_url", "varchar(1023)"},
		{"community", "users", "BASE TABLE", "id", "int(11)"},
		{"community", "users", "BASE TABLE", "season", "int(11)"},
		{"community", "users", "BASE TABLE", "user", "varchar(255)"},
		{"community", "users", "BASE TABLE", "team_id", "int(11)"},
	}
	ss.LoadSchema(schema)
	// resultSet := ss.BatchGenSQL(1)
	// for _,v:=range resultSet{
	// 	t.Log(v)
	// }
	node := ss.SelectStmt(10)

	ss.SetDB("community")
	sql, err :=	ss.Walk(node)

	for i := 0; i < 1; i++ {
		t.Log(i, ss.rdRange(1, 5))
		node := ss.SelectStmt(ss.rdRange(1, 600))
		_, _ =	ss.Walk(node)
	}

	if err != nil {
		t.Fatalf("walk error %v", err)
	}
	t.Log(sql)
}
