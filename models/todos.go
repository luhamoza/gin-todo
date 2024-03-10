package models

import "github.com/luhamoza/gin-todo/db"

type Todo struct {
	ID        int64
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todoList []Todo

func (t Todo) Save() error {
	sqlStmt := `INSERT INTO todos (title,completed) VALUES (?,?)`
	stmt, err := db.SqliteDB.Prepare(sqlStmt)
	if err != nil {
		panic("Could not prepare statement")
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(t.Title, t.Completed)
	if err != nil {
		panic("Could not execute statement")
		return err
	}
	t.ID, err = result.LastInsertId()
	return err
}

func GetAllTodos() []Todo {
	return todoList
}
