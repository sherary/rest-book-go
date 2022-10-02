package controllers

import (
	"errors"
	"net/http"
	"strconv"

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

func GetUserByID(id int) (*models.User, error) {
	for i, user := range userDatabase {
		if user.ID == id {
			return &userDatabase[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := GetUserByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
