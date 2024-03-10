package db

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var SqliteDB *sql.DB

func InitDb() {
	var err error
	SqliteDB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not initialize db")
	}
	SqliteDB.SetMaxOpenConns(25)
	SqliteDB.SetMaxIdleConns(25)
	SqliteDB.SetConnMaxLifetime(5 * time.Minute)

	createTodoTable()
}

func createTodoTable() {
	sqlStmt := `CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0)`
	_, err := SqliteDB.Exec(sqlStmt)
	if err != nil {
		panic("Could not create todos table")
	}
}
