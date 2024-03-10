package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/luhamoza/gin-todo/models"
)


func getTodos(context *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Could not get todos"})
		return
	}
	context.JSON(http.StatusOK, todos)
}
func getTodoByID(context *gin.Context) {
	id,err := strconv.ParseUint(  context.Param("id"),9,32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse id"})
		return
	}
	todo, err := models.GetTodoByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Could not get todo, todo does not exist"})
		return
	}
	context.JSON(http.StatusOK, todo)
}
func createTodos(context *gin.Context) {
	var todo models.Todo
	err := context.ShouldBindJSON(&todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse request"})
		return
	}
	err = todo.CreateTodo()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Could not create todo"})
		return
	}
	todo.ID = uuid.New().ID(); 
	context.JSON(http.StatusCreated, gin.H{"msg": "Todo created", "todo": todo})
}
