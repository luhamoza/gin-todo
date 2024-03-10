package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/luhamoza/gin-todo/db"
	"github.com/luhamoza/gin-todo/models"
)

func main() {
	db.InitDb()
	server := gin.Default()
	server.GET("/todos", getTodos)
	server.POST("/todos", createTodos)
	server.Run(":3000")
}
func getTodos(context *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Could not get todos"})
		return
	}
	context.JSON(http.StatusOK, todos)
}
func createTodos(context *gin.Context) {
	var todo models.Todo
	err := context.ShouldBindJSON(&todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse request"})
		return
	}
	err = todo.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Could not save todo"})
		return
	}
	todo.ID = uuid.New().ID(); 
	context.JSON(http.StatusCreated, gin.H{"msg": "Todo created", "todo": todo})
}
