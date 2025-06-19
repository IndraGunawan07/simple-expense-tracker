package controllers

import (
	"expense-tracker/database"
	"expense-tracker/repository"
	"expense-tracker/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllExpense(c *gin.Context) {
	var (
		result gin.H
	)

	// Retrieve the userID set by the middleware
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "User ID not found"})
		return
	}

	// Type assert to int (since we stored an int)
	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid user ID type"})
		return
	}

	expense, err := repository.GetAllExpense(database.DbConnection, userIDInt)

	if err != nil {
		result = gin.H{
			"data": err.Error(),
		}
	} else {
		result = gin.H{
			"data": expense,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertExpense(c *gin.Context) {
	var insertExpense structs.InsertExpense

	// Retrieve the userID set by the middleware
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "User ID not found"})
		return
	}

	err := c.BindJSON(&insertExpense)
	if err != nil {
		panic(err)
	}

	date, err := time.Parse("2006-01-02", insertExpense.Dates)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	expense := structs.Expense{
		UserID:      userID.(int),
		CategoryID:  insertExpense.CategoryID,
		Types:       insertExpense.Types,
		Dates:       date,
		Amount:      insertExpense.Amount,
		Description: insertExpense.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	id, err := repository.InsertExpense(database.DbConnection, expense)
	if err != nil {
		panic(err)
	}

	expense.ID = id
	c.JSON(http.StatusOK, expense)
}

func UpdateExpense(c *gin.Context) {
	var updateExpense structs.InsertExpense

	id, _ := strconv.Atoi(c.Param("id"))

	// Retrieve the userID set by the middleware
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "User ID not found"})
		return
	}

	err := c.BindJSON(&updateExpense)
	if err != nil {
		panic(err)
	}

	date, err := time.Parse("2006-01-02", updateExpense.Dates)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	expense := structs.Expense{
		ID:          id,
		UserID:      userID.(int),
		CategoryID:  updateExpense.CategoryID,
		Types:       updateExpense.Types,
		Dates:       date,
		Amount:      updateExpense.Amount,
		Description: updateExpense.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = repository.UpdateExpense(database.DbConnection, expense)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, expense)
}

func DeleteExpense(c *gin.Context) {
	var expense structs.Expense
	id, _ := strconv.Atoi(c.Param("id"))

	expense.ID = id
	err := repository.DeleteExpense(database.DbConnection, expense)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted successfully",
	})
}
