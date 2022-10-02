package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"parkee.com/library/src/models"
)

var transactionDatabase = []models.Transaction{}

func GetTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, transactionDatabase)
}

func setBorrowingStatus(id int) bool {
	var status bool
	return status
}

func CreateTransaction(c *gin.Context) {
	var input models.CreateTransactionInput
	bookID, _ := strconv.Atoi(c.Param("book_id"))
	end_date, _ := time.Parse("2006-01-02 15:04:05", c.Param("end_date"))
	status := setBorrowingStatus(bookID)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, book := range bookDatabase {
		if bookID == book.ID && book.Stock > 0 {
			status = false
		}
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	deadline := end_date.AddDate(0, 0, 7).Format("2006-01-02 15:04:05")

	transaction := models.Transaction{
		ID:        len(userDatabase) + 1,
		StartDate: now,
		EndDate:   deadline,
		BookId:    input.BookId,
	}

	transactionDatabase = append(transactionDatabase, transaction)

	if status {
		c.JSON(http.StatusOK, gin.H{"data": transaction})
		return
	}

	c.JSON(http.StatusUnavailableForLegalReasons, gin.H{"message": "cannot borrow item"})
}
