package controllers

import (
	"final-project/database"
	"final-project/models"
	"final-project/repository"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProduct(c *gin.Context) {
	var (
		result gin.H
	)

	product, err := repository.GetAllProduct(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": product,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertProduct(c *gin.Context) {
	var product models.Product

	err := c.ShouldBindJSON(&product)
	if err != nil {
		panic(err)
	}

	if product.Price <= 0 {
		panic(err)
	}
	_, err2 := url.ParseRequestURI(product.Image_url)
	if err2 != nil {
		url, _ := regexp.MatchString("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)", product.Image_url)
		if !url {
			panic(err)
		}
	}

	err = repository.InsertProduct(database.DbConnection, product)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Sukses Menambahkan Produk",
	})
}

func UpdateProduct(c *gin.Context) {
	var product models.Product

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&product)
	if err != nil {
		panic(err)
	}

	product.ID = int64(id)

	_, err2 := url.ParseRequestURI(product.Image_url)
	if err2 != nil {
		url, _ := regexp.MatchString("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)", product.Image_url)
		if !url {
			panic(err)
		}
	}

	err = repository.UpdateProduct(database.DbConnection, &product)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Sukses Update Produk",
	})
}

func DeleteProduct(c *gin.Context) {
	var product models.Product

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	product.ID = int64(id)

	err = repository.DeleteProduct(database.DbConnection, product)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Sukses Menghapus Produk",
	})
}
