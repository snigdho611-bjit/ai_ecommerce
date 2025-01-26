package main

import (
	"ecommerce/controllers"
	"ecommerce/middleware"
	"ecommerce/models"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()

	// product routes
	r.POST("/api/products", middleware.AuthMiddleware(), controllers.CreateProduct)
	r.GET("/api/products", controllers.GetAllProducts)
	r.GET("/api/products/:id", controllers.GetProductByID)
	r.PUT("/api/products/update/:id", middleware.AuthMiddleware(), controllers.UpdateProduct)
	r.DELETE("/api/products/:id", middleware.AuthMiddleware(), controllers.DeleteProductByID)
	r.GET("/api/products/filter", middleware.AuthMiddleware(), controllers.GetFilteredProducts)

	// User registration route
	r.POST("/api/users/register", controllers.RegisterUser)
	r.POST("/api/users/login", controllers.LoginUser)
	r.POST("/api/users/logout", middleware.AuthMiddleware(), controllers.LogoutUser)

	// cart routes
	r.POST("/api/cart", middleware.AuthMiddleware(), controllers.AddItemToCart)
	r.DELETE("/api/cart/:product_id", middleware.AuthMiddleware(), controllers.RemoveItemFromCart)
	r.PUT("/api/cart/:product_id", middleware.AuthMiddleware(), controllers.UpdateCartItem)
	r.GET("/api/cart", middleware.AuthMiddleware(), controllers.GetCart)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
