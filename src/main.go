package main

import (
	"log"
	"os"

	"example.com/m/src/aws"
	"example.com/m/src/controllers"
	"example.com/m/src/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

func main() {
	LoadEnv()
	session := aws.ConnectAws()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("session", session)
		c.Next()
	})

	models.ConnectDatabase() // database connection

	registerEndpoints(r)

	r.Static("/static", "../static") // expose static directory to give access to uploaded file

	r.Run(":8081") // start server at port 8081
}

func registerEndpoints(r *gin.Engine) {
	books := r.Group("/books")
	{
		books.GET("/", controllers.FindBooks)
		books.GET("/:id", controllers.FindBook)
		books.POST("/", controllers.CreateBook)
		books.PATCH("/:id", controllers.UpdateBook)
		books.DELETE("/:id", controllers.DeleteBook)
	}

	uploads := r.Group("/uploads")
	{
		uploads.POST("/", controllers.Upload)
		uploads.POST("/multiple", controllers.MultipleFileUpload)
	}
}
