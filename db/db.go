package db

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)
var sqliteDB *sql.DB
func InitDb() {
	var err error
	sqliteDB, err = sql.Open("sqlite3","api.db")
	if err != nil {
		panic("Could not initialize db")
	}
	sqliteDB.SetMaxOpenConns(25)
	sqliteDB.SetMaxIdleConns(25)
	sqliteDB.SetConnMaxLifetime(5*time.Minute)

	createTodoTable()
}
func createTodoTable() {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0
	)`
	_, err := sqliteDB.Exec(sqlStmt)
	if err != nil {
		panic("Could not create todos table")
	}
}