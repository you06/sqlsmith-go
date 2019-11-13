package sqlsmith

import "testing"

var testSchema = [][5]string{
	// {"community", "comments", "BASE TABLE", "id", "int(11)"},
	// {"community", "comments", "BASE TABLE", "owner", "varchar(255)"},
	// {"community", "comments", "BASE TABLE", "repo", "varchar(255)"},
	// {"community", "comments", "BASE TABLE", "comment_id", "int(11)"},
	// {"community", "comments", "BASE TABLE", "comment_type", "varchar(128)"},
	// {"community", "comments", "BASE TABLE", "pull_number", "int(11)"},
	// {"community", "comments", "BASE TABLE", "body", "text"},
	// {"community", "comments", "BASE TABLE", "user", "varchar(255)"},
	// {"community", "comments", "BASE TABLE", "url", "varchar(1023)"},
	// {"community", "comments", "BASE TABLE", "association", "varchar(255)"},
	// {"community", "comments", "BASE TABLE", "relation", "varchar(255)"},
	// {"community", "comments", "BASE TABLE", "created_at", "timestamp"},
	// {"community", "comments", "BASE TABLE", "updated_at", "timestamp"},
	// {"community", "picks", "BASE TABLE", "id", "int(11)"},
	// {"community", "picks", "BASE TABLE", "season", "int(11)"},
	// {"community", "picks", "BASE TABLE", "task_id", "int(11)"},
	// {"community", "picks", "BASE TABLE", "teamID", "int(11)"},
	// {"community", "picks", "BASE TABLE", "user", "varchar(255)"},
	// {"community", "picks", "BASE TABLE", "pull_number", "int(11)"},
	// {"community", "picks", "BASE TABLE", "status", "varchar(128)"},
	// {"community", "picks", "BASE TABLE", "created_at", "timestamp"},
	// {"community", "picks", "BASE TABLE", "updated_at", "timestamp"},
	// {"community", "picks", "BASE TABLE", "closed_at", "timestamp"},
	// {"community", "pulls", "BASE TABLE", "id", "int(11)"},
	// {"community", "pulls", "BASE TABLE", "owner", "varchar(255)"},
	// {"community", "pulls", "BASE TABLE", "repo", "varchar(255)"},
	// {"community", "pulls", "BASE TABLE", "pull_number", "int(11)"},
	// {"community", "pulls", "BASE TABLE", "title", "text"},
	// {"community", "pulls", "BASE TABLE", "body", "text"},
	// {"community", "pulls", "BASE TABLE", "user", "varchar(255)"},
	// {"community", "pulls", "BASE TABLE", "association", "varchar(255)"},
	// {"community", "pulls", "BASE TABLE", "relation", "varchar(255)"},
	// {"community", "pulls", "BASE TABLE", "label", "text"},
	// {"community", "pulls", "BASE TABLE", "status", "varchar(128)"},
	// {"community", "pulls", "BASE TABLE", "created_at", "timestamp"},
	// {"community", "pulls", "BASE TABLE", "updated_at", "timestamp"},
	// {"community", "pulls", "BASE TABLE", "closed_at", "timestamp"},
	// {"community", "pulls", "BASE TABLE", "merged_at", "timestamp"},
	// {"community", "tasks", "BASE TABLE", "id", "int(11)"},
	// {"community", "tasks", "BASE TABLE", "season", "int(11)"},
	// {"community", "tasks", "BASE TABLE", "complete_user", "varchar(255)"},
	// {"community", "tasks", "BASE TABLE", "complete_team", "int(11)"},
	// {"community", "tasks", "BASE TABLE", "owner", "varchar(255)"},
	// {"community", "tasks", "BASE TABLE", "repo", "varchar(255)"},
	// {"community", "tasks", "BASE TABLE", "title", "varchar(2047)"},
	// {"community", "tasks", "BASE TABLE", "issue_number", "int(11)"},
	// {"community", "tasks", "BASE TABLE", "pull_number", "int(11)"},
	// {"community", "tasks", "BASE TABLE", "level", "varchar(255)"},
	// {"community", "tasks", "BASE TABLE", "min_score", "int(11)"},
	// {"community", "tasks", "BASE TABLE", "score", "int(11)"},
	// {"community", "tasks", "BASE TABLE", "status", "varchar(255)"},
	// {"community", "tasks", "BASE TABLE", "created_at", "timestamp"},
	// {"community", "tasks", "BASE TABLE", "expired", "varchar(255)"},
	// {"community", "teams", "BASE TABLE", "id", "int(11)"},
	// {"community", "teams", "BASE TABLE", "season", "int(11)"},
	// {"community", "teams", "BASE TABLE", "name", "varchar(255)"},
	// {"community", "teams", "BASE TABLE", "issue_url", "varchar(1023)"},
	// {"community", "users", "BASE TABLE", "id", "int(11)"},
	// {"community", "users", "BASE TABLE", "season", "int(11)"},
	// {"community", "users", "BASE TABLE", "user", "varchar(255)"},
	// {"community", "users", "BASE TABLE", "team_id", "int(11)"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "kjzecwur", "text"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "uogdcz", "text"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "muqffml", "timestamp"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "xcnotemvj", "timestamp"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "zmfer", "varchar(539)"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "kcgdb", "int(8)"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "tnalhqynf", "int(14)"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "id", "int(11)"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "imedrgbz", "varchar(467)"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "vpsxee", "int(10)"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "fpznqz", "timestamp"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "gnopbfu", "int(8)"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "xbsucfyrv", "int(11)"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "axqwthgdt", "timestamp"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "oznnhnqa", "int(14)"},
	{"sqlsmith", "tfbilhei", "BASE TABLE", "fumik", "varchar(386)"},
	{"sqlsmith", "jirllla", "BASE TABLE", "id", "int(11)"},
	{"sqlsmith", "jirllla", "BASE TABLE", "rqetq", "text"},
	{"sqlsmith", "jirllla", "BASE TABLE", "arjumsu", "timestamp"},
	{"sqlsmith", "jirllla", "BASE TABLE", "ktejdzpb", "varchar(679)"},
	{"sqlsmith", "jirllla", "BASE TABLE", "ilbgyhi", "int(14)"},
	{"sqlsmith", "jmydrjyww", "BASE TABLE", "dokvhkep", "varchar(397)"},
	{"sqlsmith", "jmydrjyww", "BASE TABLE", "zmnupz", "text"},
	{"sqlsmith", "jmydrjyww", "BASE TABLE", "eddpy", "int(16)"},
	{"sqlsmith", "jmydrjyww", "BASE TABLE", "nufhhor", "int(8)"},
	{"sqlsmith", "jmydrjyww", "BASE TABLE", "okccbb", "varchar(1856)"},
	{"sqlsmith", "jmydrjyww", "BASE TABLE", "akjfaeod", "timestamp"},
	{"sqlsmith", "jmydrjyww", "BASE TABLE", "mcizyqs", "int(11)"},
	{"sqlsmith", "jmydrjyww", "BASE TABLE", "uzmidqka", "int(12)"},
	{"sqlsmith", "jmydrjyww", "BASE TABLE", "id", "int(11)"},
	{"sqlsmith", "jmydrjyww", "BASE TABLE", "skbxq", "varchar(320)"},
	{"sqlsmith", "jmydrjyww", "BASE TABLE", "adqvzs", "text"},
	{"sqlsmith", "jmydrjyww", "BASE TABLE", "ugmpyr", "timestamp"},
	{"sqlsmith", "jmydrjyww", "BASE TABLE", "olzogwctg", "varchar(1675)"},
	{"sqlsmith", "grhuztuct", "BASE TABLE", "gisclis", "timestamp"},
	{"sqlsmith", "grhuztuct", "BASE TABLE", "jwjsx", "timestamp"},
	{"sqlsmith", "grhuztuct", "BASE TABLE", "xuqbap", "text"},
	{"sqlsmith", "grhuztuct", "BASE TABLE", "srtsqgh", "timestamp"},
	{"sqlsmith", "grhuztuct", "BASE TABLE", "id", "int(11)"},
	{"sqlsmith", "grhuztuct", "BASE TABLE", "zwyhiysp", "int(19)"},
	{"sqlsmith", "grhuztuct", "BASE TABLE", "tbwdudrcx", "int(11)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "id", "int(11)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "gdldulhl", "int(11)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "pumhir", "int(9)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "swijx", "int(19)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "hgsrbxg", "varchar(1354)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "sxiujciw", "int(18)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "pyhzeti", "timestamp"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "kalfpitti", "varchar(1698)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "jqhfwcib", "int(14)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "qvctr", "int(17)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "ekglwupgv", "text"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "pmeucn", "varchar(870)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "ydyutnwq", "varchar(702)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "alvna", "int(11)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "mmwlyp", "timestamp"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "xofcxgnfi", "timestamp"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "eaezsn", "int(18)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "dxuvctj", "varchar(1851)"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "wmbowsyfg", "text"},
	{"sqlsmith", "jpwljcv", "BASE TABLE", "dyshaw", "varchar(863)"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "sakmtx", "timestamp"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "prvzn", "int(15)"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "vsrrczojy", "text"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "tjlcvsbgn", "int(16)"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "xygqtm", "text"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "utjurwu", "timestamp"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "vrrmc", "text"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "id", "int(11)"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "yicgw", "timestamp"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "xxidplox", "int(13)"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "txhmr", "int(18)"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "cotyu", "int(16)"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "othulefk", "int(18)"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "ypzqqnght", "timestamp"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "jsqnh", "int(17)"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "wrtnhpsmx", "text"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "nwoqr", "int(10)"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "htlkfkg", "text"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "fihlh", "int(12)"},
	{"sqlsmith", "zpyqqbeh", "BASE TABLE", "nyweuiou", "varchar(1092)"},
	{"sqlsmith", "jzdmilto", "BASE TABLE", "swqff", "text"},
	{"sqlsmith", "jzdmilto", "BASE TABLE", "sfyqnpko", "int(15)"},
	{"sqlsmith", "jzdmilto", "BASE TABLE", "eamdhqdqr", "varchar(281)"},
	{"sqlsmith", "jzdmilto", "BASE TABLE", "wnogc", "text"},
	{"sqlsmith", "jzdmilto", "BASE TABLE", "irlfae", "int(18)"},
	{"sqlsmith", "jzdmilto", "BASE TABLE", "id", "int(11)"},
	{"sqlsmith", "jzdmilto", "BASE TABLE", "unvxgveab", "varchar(2035)"},
	{"sqlsmith", "jzdmilto", "BASE TABLE", "xiafkd", "varchar(917)"},
	{"sqlsmith", "jzdmilto", "BASE TABLE", "ubmbfrk", "timestamp"},
	{"sqlsmith", "jzdmilto", "BASE TABLE", "mitnzlh", "text"},
	{"sqlsmith", "jzdmilto", "BASE TABLE", "mbwvfzeu", "timestamp"},
	{"sqlsmith", "jzdmilto", "BASE TABLE", "kwetu", "timestamp"},
	{"sqlsmith", "jzdmilto", "BASE TABLE", "wnhhtgkn", "timestamp"},
	{"sqlsmith", "dfjeuh", "BASE TABLE", "id", "int(11)"},
	{"sqlsmith", "dfjeuh", "BASE TABLE", "pimajgeb", "varchar(347)"},
	{"sqlsmith", "dfjeuh", "BASE TABLE", "srysvmhw", "timestamp"},
	{"sqlsmith", "dfjeuh", "BASE TABLE", "wootcbi", "int(14)"},
	{"sqlsmith", "dfjeuh", "BASE TABLE", "imitlr", "text"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "ufezz", "int(17)"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "zwepsruk", "int(12)"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "hzbrzjxi", "timestamp"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "kgbuaiaf", "int(18)"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "ydkvkxcy", "timestamp"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "uwkuyfjhj", "varchar(1048)"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "liqtwvuof", "timestamp"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "jmxjbuy", "int(12)"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "qnwekvg", "timestamp"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "bxarcwl", "int(16)"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "xbsjhgvnt", "varchar(1863)"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "jixkqeexc", "int(13)"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "icnuxjzfr", "timestamp"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "id", "int(11)"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "ajsckeqr", "varchar(1272)"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "lnifan", "timestamp"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "bonlvddv", "int(16)"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "zhmyxdou", "text"},
	{"sqlsmith", "kglyaxwu", "BASE TABLE", "vrhowqvlk", "text"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "uezlhb", "varchar(633)"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "xhqxbww", "varchar(736)"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "iurxcoin", "int(17)"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "mhpmtcb", "timestamp"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "mmhplf", "int(14)"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "id", "int(11)"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "gfdvw", "int(17)"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "uaqxtojd", "timestamp"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "uagpdj", "varchar(1966)"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "pegutjhmy", "int(10)"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "kjnih", "text"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "xgdhicw", "text"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "tjmjwthmj", "varchar(675)"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "qookk", "int(10)"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "crjcdk", "int(11)"},
	{"sqlsmith", "wviwbof", "BASE TABLE", "idikrkqul", "int(17)"},
}

var testSchemaName = "sqlsmith"

func TestSQLSmith_Select(t *testing.T) {
	ss := New()
	ss.LoadSchema(testSchema)

	ss.SetDB(testSchemaName)
	sql, _ := ss.SelectStmt(3)
	t.Log(sql)
}

func TestSQLSmith_Update(t *testing.T) {
	ss := New()
	ss.LoadSchema(testSchema)

	ss.SetDB(testSchemaName)

	for i := 0; i < 100000000; i++ {
		sql, err := ss.UpdateStmt()
		if err != nil {
			t.Log(sql, err)
		}
	}
	sql, _ := ss.UpdateStmt()
	t.Log(sql)
}

func TestSQLSmith_Insert(t *testing.T) {
	ss := New()
	ss.LoadSchema(testSchema)

	ss.SetDB(testSchemaName)

	for i := 0; i < 1000; i++ {
		sql, err := ss.InsertStmtAST()
		if err != nil {
			t.Log(sql, err)
		}
	}
	sql, err := ss.InsertStmtAST()
	t.Log(sql, err)
}
