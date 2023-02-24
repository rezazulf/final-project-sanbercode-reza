package controllers

import (
	"final-project/database"
	"final-project/middleware"
	"final-project/models"
	"final-project/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	var (
		result gin.H
	)

	order, err := repository.GetAllTransactions(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": order,
		}
	}

	c.JSON(http.StatusOK, result)
}

func PostTransactions(c *gin.Context) {
	// authenticate session
	middleware.Auth("Customer")

	customer_id, _ := strconv.Atoi(c.Param("id"))
	item_id, _ := strconv.Atoi(c.Param("item_id"))

	var (
		err          error
		result       string
		dataCustomer models.Users
		dataProduct  models.Product
		dataCheck    models.Transactions
		dataCheckout *models.Transactions
		isTrue       bool
	)

	err = c.ShouldBindJSON(&dataCheck) // baca banyak item dibeli
	if err != nil {
		panic(err)
	}

	dataCustomer.ID = int64(customer_id)
	dataProduct.ID = int64(item_id)

	err, result, dataCheckout, isTrue = repository.PostTransactions(database.DbConnection, &dataCustomer, &dataProduct, &dataCheck)
	if err != nil {
		panic(err)
	}

	if isTrue {
		c.JSON(http.StatusOK, gin.H{
			"message":     result,
			"user_name":   dataCustomer.Username,
			"item_bought": dataProduct.Name,
			"sum_item":    dataCheckout.Sum_item,
			"bill":        dataCheckout.Payment_bills,
			"balance":     dataCustomer.Balance,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result,
		})
	}

}
