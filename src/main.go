package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"parkee.com/library/src/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Bye World!"})
	})

	//Book routes
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBook)
	router.POST("/books/create", controllers.CreateBook)

	// User Routes
	router.GET("/users", controllers.GetUsers)
	router.POST("/users/create", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUser)

	router.Run("localhost:3000")
}
