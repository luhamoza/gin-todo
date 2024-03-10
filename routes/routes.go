package routes

import "github.com/gin-gonic/gin"


func RegisterRoutes(server *gin.Engine){
	server.GET("/todos", getTodos)
	server.GET("/todos/:id",getTodoByID)
	server.POST("/todos", createTodos)
	// server.PUT("/todos/:id", updateTodo)
	// server.DELETE("/todos/:id", deleteTodo)
}