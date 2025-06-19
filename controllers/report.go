package controllers

import (
	"expense-tracker/database"
	"expense-tracker/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetReport(c *gin.Context) {
	var (
		result gin.H
	)

	// Initialize filters
	filters := map[string]interface{}{}

	// Parse optional start_date
	if startDateStr := c.Query("start_date"); startDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format. Use YYYY-MM-DD"})
			return
		}
		filters["start_date"] = startDate
	}

	// Parse optional end_date
	if endDateStr := c.Query("end_date"); endDateStr != "" {
		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format. Use YYYY-MM-DD"})
			return
		}
		filters["end_date"] = endDate
	}

	// Validate date range if both dates are provided
	if startDate, ok1 := filters["start_date"].(time.Time); ok1 {
		if endDate, ok2 := filters["end_date"].(time.Time); ok2 {
			if endDate.Before(startDate) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "end_date must be after start_date"})
				return
			}
		}
	}

	// Parse optional type
	if typeStr := c.Query("type"); typeStr != "" {
		typeInt, err := strconv.Atoi(typeStr)
		if err != nil || (typeInt != 1 && typeInt != 2) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "type must be either 1 or 2"})
			return
		}
		filters["type"] = typeInt
	}

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

	// Query database with applicable filters
	// items, err := repository.GetReport(database.DbConnection, filters)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"filters": filters,
	// 	"items":   items,
	// })

	expense, grandTotal, err := repository.GetReport(database.DbConnection, userIDInt, filters)

	if err != nil {
		result = gin.H{
			"data": err.Error(),
		}
	} else {
		result = gin.H{
			"data":        expense,
			"grand_total": grandTotal,
		}
	}

	c.JSON(http.StatusOK, result)
}
