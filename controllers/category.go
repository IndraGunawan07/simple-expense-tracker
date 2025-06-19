package controllers

import (
	"expense-tracker/database"
	"expense-tracker/repository"
	"expense-tracker/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
	var (
		result gin.H
	)

	category, err := repository.GetAllCategory(database.DbConnection)

	if err != nil {
		result = gin.H{
			"data": err.Error(),
		}
	} else {
		result = gin.H{
			"data": category,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var category structs.Category

	err := c.BindJSON(&category)
	if err != nil {
		panic(err)
	}

	id, err := repository.InsertCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	category.ID = id
	c.JSON(http.StatusOK, category)
}

func UpdateCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.ID = id

	err = repository.UpdateCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	category.ID = id
	err := repository.DeleteCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted successfully",
	})
}
