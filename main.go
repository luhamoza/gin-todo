package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()
	server.GET("/todos",getTodos)
	server.Run(":3000")
}
func getTodos(context *gin.Context){
	context.JSON(http.StatusOK,gin.H{"msg":"Hello World"})
}
