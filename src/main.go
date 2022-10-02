package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"parkee.com/library/src/controllers"
)

// func getBooks(context *gin.Context) {
// 	context.IndentedJSON(http.StatusOK, books)
// }

// func addBook(context *gin.Context) {
// 	var book models.Book

// 	if err := context.BindJSON(&book); err != nil {
// 		return
// 	}

// 	books = append(books, book)
// 	context.IndentedJSON(http.StatusCreated, book)
// }

// func getBook(context *gin.Context) {
// 	id := context.Params("id")
// 	book, err := getBookByID(id)

// 	if err != nil {
// 		context.IndentedJSON(http.StatusNotFound, gin.H{})
// 	}

// }

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Bye World!"})
	})

	//Book routes
	router.GET("/books", controllers.GetBooks)
	// router.GET("/book/:id", getBook)
	router.POST("/books/create", controllers.CreateBook)

	// User Routes
	router.GET("/users", controllers.GetUsers)
	router.POST("/users/create", controllers.CreateUser)

	router.Run("localhost:3000")
}
