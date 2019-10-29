package sqlsmith

import "testing"

func TestSQLSmith_ToSQL(t *testing.T) {
	ss := New()
	ss.LoadSchema([][5]string{
		[5]string{"community", "comments", "BASE TABLE", "id", "int(11)"},
		[5]string{"community", "comments", "BASE TABLE", "owner", "varchar(255)"},
		[5]string{"community", "comments", "BASE TABLE", "repo", "varchar(255)"},
		[5]string{"community", "comments", "BASE TABLE", "comment_id", "int(11)"},
		[5]string{"community", "comments", "BASE TABLE", "comment_type", "varchar(128)"},
		[5]string{"community", "comments", "BASE TABLE", "pull_number", "int(11)"},
		[5]string{"community", "comments", "BASE TABLE", "body", "text"},
		[5]string{"community", "comments", "BASE TABLE", "user", "varchar(255)"},
		[5]string{"community", "comments", "BASE TABLE", "url", "varchar(1023)"},
		[5]string{"community", "comments", "BASE TABLE", "association", "varchar(255)"},
		[5]string{"community", "comments", "BASE TABLE", "relation", "varchar(255)"},
		[5]string{"community", "comments", "BASE TABLE", "created_at", "timestamp"},
		[5]string{"community", "comments", "BASE TABLE", "updated_at", "timestamp"},
		[5]string{"community", "picks", "BASE TABLE", "id", "int(11)"},
		[5]string{"community", "picks", "BASE TABLE", "season", "int(11)"},
		[5]string{"community", "picks", "BASE TABLE", "task_id", "int(11)"},
		[5]string{"community", "picks", "BASE TABLE", "teamID", "int(11)"},
		[5]string{"community", "picks", "BASE TABLE", "user", "varchar(255)"},
		[5]string{"community", "picks", "BASE TABLE", "pull_number", "int(11)"},
		[5]string{"community", "picks", "BASE TABLE", "status", "varchar(128)"},
		[5]string{"community", "picks", "BASE TABLE", "created_at", "timestamp"},
		[5]string{"community", "picks", "BASE TABLE", "updated_at", "timestamp"},
		[5]string{"community", "picks", "BASE TABLE", "closed_at", "timestamp"},
		[5]string{"community", "pulls", "BASE TABLE", "id", "int(11)"},
		[5]string{"community", "pulls", "BASE TABLE", "owner", "varchar(255)"},
		[5]string{"community", "pulls", "BASE TABLE", "repo", "varchar(255)"},
		[5]string{"community", "pulls", "BASE TABLE", "pull_number", "int(11)"},
		[5]string{"community", "pulls", "BASE TABLE", "title", "text"},
		[5]string{"community", "pulls", "BASE TABLE", "body", "text"},
		[5]string{"community", "pulls", "BASE TABLE", "user", "varchar(255)"},
		[5]string{"community", "pulls", "BASE TABLE", "association", "varchar(255)"},
		[5]string{"community", "pulls", "BASE TABLE", "relation", "varchar(255)"},
		[5]string{"community", "pulls", "BASE TABLE", "label", "text"},
		[5]string{"community", "pulls", "BASE TABLE", "status", "varchar(128)"},
		[5]string{"community", "pulls", "BASE TABLE", "created_at", "timestamp"},
		[5]string{"community", "pulls", "BASE TABLE", "updated_at", "timestamp"},
		[5]string{"community", "pulls", "BASE TABLE", "closed_at", "timestamp"},
		[5]string{"community", "pulls", "BASE TABLE", "merged_at", "timestamp"},
		[5]string{"community", "tasks", "BASE TABLE", "id", "int(11)"},
		[5]string{"community", "tasks", "BASE TABLE", "season", "int(11)"},
		[5]string{"community", "tasks", "BASE TABLE", "complete_user", "varchar(255)"},
		[5]string{"community", "tasks", "BASE TABLE", "complete_team", "int(11)"},
		[5]string{"community", "tasks", "BASE TABLE", "owner", "varchar(255)"},
		[5]string{"community", "tasks", "BASE TABLE", "repo", "varchar(255)"},
		[5]string{"community", "tasks", "BASE TABLE", "title", "varchar(2047)"},
		[5]string{"community", "tasks", "BASE TABLE", "issue_number", "int(11)"},
		[5]string{"community", "tasks", "BASE TABLE", "pull_number", "int(11)"},
		[5]string{"community", "tasks", "BASE TABLE", "level", "varchar(255)"},
		[5]string{"community", "tasks", "BASE TABLE", "min_score", "int(11)"},
		[5]string{"community", "tasks", "BASE TABLE", "score", "int(11)"},
		[5]string{"community", "tasks", "BASE TABLE", "status", "varchar(255)"},
		[5]string{"community", "tasks", "BASE TABLE", "created_at", "timestamp"},
		[5]string{"community", "tasks", "BASE TABLE", "expired", "varchar(255)"},
		[5]string{"community", "teams", "BASE TABLE", "id", "int(11)"},
		[5]string{"community", "teams", "BASE TABLE", "season", "int(11)"},
		[5]string{"community", "teams", "BASE TABLE", "name", "varchar(255)"},
		[5]string{"community", "teams", "BASE TABLE", "issue_url", "varchar(1023)"},
		[5]string{"community", "users", "BASE TABLE", "id", "int(11)"},
		[5]string{"community", "users", "BASE TABLE", "season", "int(11)"},
		[5]string{"community", "users", "BASE TABLE", "user", "varchar(255)"},
		[5]string{"community", "users", "BASE TABLE", "team_id", "int(11)"},
	})
	// resultSet := ss.BatchGenSQL(1)
	// for _,v:=range resultSet{
	// 	t.Log(v)
	// }
	node := ss.SelectStmt(5)

	ss.SetDB("community")
	sql, err :=	ss.Walk(node)

	if err != nil {
		t.Fatalf("walk error %v", err)
	}
	t.Log(sql)
}
