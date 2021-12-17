package main

import (
	"example.com/m/src/controllers"
	"example.com/m/src/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase() // new

	registerEndpoints(r)

	r.Run()
}

func registerEndpoints(r *gin.Engine) {
	// books
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
}
