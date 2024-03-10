package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/luhamoza/gin-todo/db"
	"github.com/luhamoza/gin-todo/routes"
)

func main() {
	db.InitDb()
	server := gin.Default()
	fmt.Println("here")
	routes.RegisterRoutes(server)
	server.Run(":3000")
}