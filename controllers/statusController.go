package controllers

import (
	"final-project/database"
	"final-project/models"
	"final-project/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllStatus(c *gin.Context) {
	var (
		result gin.H
	)

	status, err := repository.GetAllStatus(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": status,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertStatus(c *gin.Context) {
	var status models.Status

	err := c.ShouldBindJSON(&status)
	if err != nil {
		panic(err)
	}

	err = repository.InsertStatus(database.DbConnection, status)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Sukses Menambahkan Status",
	})
}
