package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luhamoza/gin-todo/db"
	"github.com/luhamoza/gin-todo/middleware"
	"github.com/luhamoza/gin-todo/routes"
)

func main() {
	db.InitDb()

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	server.GET("/todos", routes.GetTodos)
	server.GET("/todos/:id",routes.GetTodoByID)
	server.POST("/todos", routes.CreateTodos)
	server.PUT("/todos/:id", routes.UpdateTodoByID)
	server.DELETE("/todos/:id",routes.DeleteTodoByID)
	
	server.Run(":3001")
}

