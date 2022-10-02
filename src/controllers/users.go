package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"parkee.com/library/src/models"
)

var userDatabase = []models.User{}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userDatabase)
}

func CreateUser(c *gin.Context) {
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{
		ID:    len(userDatabase) + 1,
		Name:  input.Name,
		NIK:   input.NIK,
		Email: input.Email,
	}

	userDatabase = append(userDatabase, user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
