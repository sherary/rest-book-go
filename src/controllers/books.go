package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"parkee.com/library/src/models"
)

var bookDatabase = []models.Book{}

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
	bookDatabase = append(bookDatabase, book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func GetBookByID(id int) (*models.Book, error) {
	for i, book := range bookDatabase {
		if book.ID == id {
			return &bookDatabase[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := GetBookByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}
