package controllers

import (
	"net/http"

	"example.com/m/src/dto"
	"example.com/m/src/models"
	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var input dto.CreateBookInput
	if error := c.ShouldBindJSON(&input); error != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": error.Error(),
			},
		)
		return
	}

	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	models.DB.Create(&book)
	c.JSON(
		http.StatusOK,
		gin.H{
			"data": book,
		},
	)
}

func FindBook(c *gin.Context) { // Get model if exist
	book, err := findBookById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	book, err := findBookById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input dto.UpdateBookInput
	if error := c.ShouldBindJSON(&input); error != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": error.Error(),
			},
		)
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	book, err := findBookById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(
		http.StatusOK,
		gin.H{"data": true},
	)
}

func findBookById(id string) (models.Book, error) {
	var book models.Book

	if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}
