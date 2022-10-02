package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"parkee.com/library/src/models"
)

var bookDatabase = []models.Book{
	// {
	// 	ID:    1,
	// 	Title: "Harry Potter",
	// 	Stock: 10,
	// 	ISBN:  "123-456-789",
	// }, {
	// 	ID:    2,
	// 	Title: "Harry Potter 2",
	// 	Stock: 12,
	// 	ISBN:  "234-567-890",
	// }, {
	// 	ID:    3,
	// 	Title: "Harry Potter 3",
	// 	Stock: 8,
	// 	ISBN:  "345-678-901",
	// },
}

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, bookDatabase)
}

func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	book := models.Book{
		ID:    len(bookDatabase) + 1,
		Title: input.Title,
		Stock: input.Stock,
		ISBN:  input.ISBN,
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}
