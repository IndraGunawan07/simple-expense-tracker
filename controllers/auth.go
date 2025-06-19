package controllers

import (
	"expense-tracker/database"
	"expense-tracker/repository"
	"expense-tracker/structs"
	"expense-tracker/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
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

	user, err := repository.GetUser(database.DbConnection, userIDInt)

	if err != nil {
		result = gin.H{
			"data": err.Error(),
		}
	} else {
		result = gin.H{
			"data": gin.H{
				"id":    user.ID,
				"email": user.Email,
			},
		}
	}

	c.JSON(http.StatusOK, result)
}

func Register(c *gin.Context) {
	var user structs.User

	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}
	id, err := repository.Register(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	user.ID = id
	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}

func Login(c *gin.Context) {
	var credentials structs.Credentials

	err := c.BindJSON(&credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request format",
		})
		return
	}

	token, err := repository.Login(database.DbConnection, credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
// 			c.Abort()
// 			return
// 		}

// 		userID, err := utils.VerifyToken(tokenString)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("userID", userID)
// 		c.Next()
// 	}
// }

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Extract the token from the Bearer string
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // means no Bearer prefix found
			c.JSON(401, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		userID, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{
				"error":   "Invalid token",
				"details": err.Error(), // Include specific error
			})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
