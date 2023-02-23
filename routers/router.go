package routers

import (
	"final-project/controllers"
	"final-project/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "low code, error handling not finished yet.",
		})
	})

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/logout", middleware.Auth("isLogin"), controllers.Logout)

	// admin area
	router.GET("/admin/home", middleware.Auth("Admin"), controllers.Validate)
	router.GET("/admin/list-admin", middleware.Auth("Admin"), controllers.GetAdmins)
	router.GET("/admin/list-user", middleware.Auth("Admin"), controllers.GetUsers)
	router.PUT("/admin/new-password", middleware.Auth("Admin"), controllers.NewPassword)
	router.GET("/admin/customer-order", middleware.Auth("Admin"), controllers.GetAllTransactions)
	router.PUT("/admin/balance/:id", middleware.Auth("Admin"), controllers.EditBalance)

	router.GET("/categories", controllers.GetAllCategory)
	router.POST("/categories", middleware.Auth("Admin"), controllers.InsertCategory)
	router.PUT("/categories/:id", middleware.Auth("Admin"), controllers.UpdateCategory)
	router.DELETE("/categories/:id", middleware.Auth("Admin"), controllers.DeleteCategory)

	router.GET("/product", controllers.GetAllProduct)
	router.POST("/product", middleware.Auth("Admin"), controllers.InsertProduct)
	router.PUT("/product/:id", middleware.Auth("Admin"), controllers.UpdateProduct)
	router.DELETE("/product/:id", middleware.Auth("Admin"), controllers.DeleteProduct)

	// user area
	router.GET("/home", middleware.Auth("Customer"), controllers.Validate)
	router.PUT("/new-password", middleware.Auth("Customer"), controllers.NewPassword)
	router.GET("/list-user", middleware.Auth("Customer"), controllers.GetUsers)
	router.GET("/customer/customer-order", middleware.Auth("Customer"), controllers.GetAllTransactions)

	router.GET("/status", controllers.GetAllStatus)
	router.POST("/status", controllers.InsertStatus)

	router.POST("/buy/:item_id/customer/:id", middleware.Auth("Customer"), controllers.PostTransactions) // membeli barang
	// router.POST("/purchase", controllers.InsertOrder) // membeli barang

	return router
}
