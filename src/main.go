package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"example.com/m/src/controllers"
	"example.com/m/src/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase() // database connection

	registerEndpoints(r)

	r.Static("/static", "../static") // expose static folder to give access to saved file

	r.Run(":8081") // start server at port 8081
}

func fileUpload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files[]"]

	fileNames := make([]string, len(files))
	for ind, file := range files {
		log.Println(ind, file.Filename)
		random := strconv.Itoa(time.Now().Local().Nanosecond())
		destination, _ := os.Create("../static/" + random + "-" + file.Filename)

		err := c.SaveUploadedFile(file, destination.Name())
		if err != nil {
			log.Println("File: " + destination.Name() + " could not be saved for")
			fileNames[ind] = ""
		} else {
			fileNames[ind] = destination.Name()
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": fileNames})
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
		uploads.POST("/", fileUpload)
	}
}
