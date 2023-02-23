package controllers

import (
	"final-project/database"
	"final-project/models"
	"final-project/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	categoryInput struct {
		Name string `json:"name"`
	}
)

func GetAllCategory(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetAllCategory(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var category models.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Category",
	})
}

func UpdateCategory(c *gin.Context) {
	var category models.Category

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.ID = int64(id)

	err = repository.UpdateCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Category",
	})
}

func DeleteCategory(c *gin.Context) {
	var category models.Category

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	category.ID = int64(id)

	err = repository.DeleteCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Category",
	})
}
