package sqlsmith

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DBConn struct {
	db *sql.DB
}

type SchemaMeta map[string]map[string]string

func NewDB(host string) (*DBConn, error){
	conn, err := sql.Open("mysql", host)
	if err!=nil{
		return nil, err
	}
	return &DBConn{conn}, nil
}

func (db *DBConn)FetchSchemaMetas(ctx context.Context) (SchemaMeta, error){
	return nil, nil
}