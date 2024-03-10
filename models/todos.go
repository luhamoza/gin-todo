package models

type Todo struct {
	ID          int
	Title       string 
	Status      bool
}

var todoList = []Todo{}

func (t Todo) Save() {
	todoList = append(todoList,t)
} 

func GetAllTodos() []Todo {
	return todoList
}
