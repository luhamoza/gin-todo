package models

import (
	"fmt"

	"github.com/luhamoza/gin-todo/db"
)
type Todo struct {
	ID       uint32 `json:"id"` 
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t Todo) CreateTodo() error  {
	sqlStmt := `INSERT INTO todos (title,completed) VALUES (?,?)`
	_,err := db.SqliteDB.Exec(sqlStmt,t.Title, t.Completed)
	if err != nil {
		return fmt.Errorf("could not execute todo: %v", err)
	}
	return err	
}

func GetAllTodos() ([]Todo, error) {
	rows, err := db.SqliteDB.Query("SELECT id,title,completed FROM todos")
	if err != nil {
		return nil, fmt.Errorf("could not get todos: %v", err)
	}
	var todoList []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan((&todo.ID),(&todo.Title),(&todo.Completed))
		if err != nil {
			return nil, fmt.Errorf("could not get todos: %v", err)
		}
		todoList = append(todoList, todo)
	}
	return todoList, nil
}
func GetTodoByID(id uint64) (Todo, error) {
	var todo Todo
	err := db.SqliteDB.QueryRow("SELECT id,title,completed FROM todos WHERE id = ?", id).Scan((&todo.ID),(&todo.Title),(&todo.Completed))
	if err != nil {
		return todo, fmt.Errorf("could not get todo: %v", err)
	}
	return todo, nil
}