package controllers

import (
	"log"
	"net/http"

	"ecommerce/models"

	"github.com/gin-gonic/gin"
)

// AddItemToCartInput represents the input for adding an item to the cart
type AddItemToCartInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

// UpdateCartItemInput represents the input for updating an item in the cart
type UpdateCartItemInput struct {
	Quantity int `json:"quantity" binding:"required"`
}

// AddItemToCart adds an item to the user's cart
func AddItemToCart(c *gin.Context) {
	log.Println("Request received for adding an item to the cart")
	var input AddItemToCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "success": false, "error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	cartItem := models.Cart{
		UserID:    userID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
	}

	if err := models.DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add item to cart", "success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart successfully", "success": true, "data": cartItem})
}

// RemoveItemFromCart removes an item from the user's cart
func RemoveItemFromCart(c *gin.Context) {
	log.Println("Request received for removing an item from the cart")
	userID := c.GetUint("user_id")
	productID := c.Param("product_id")

	var cartItem models.Cart
	if err := models.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cartItem).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Item not found in cart", "success": false, "error": err.Error()})
		return
	}

	if err := models.DB.Delete(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to remove item from cart", "success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart successfully", "success": true})
}

// UpdateCartItem updates the quantity of an item in the user's cart
func UpdateCartItem(c *gin.Context) {
	log.Println("Request received for updating an item in the cart")
	userID := c.GetUint("user_id")
	productID := c.Param("product_id")

	var input UpdateCartItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "success": false, "error": err.Error()})
		return
	}

	var cartItem models.Cart
	if err := models.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cartItem).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Item not found in cart", "success": false, "error": err.Error()})
		return
	}

	cartItem.Quantity = input.Quantity
	if err := models.DB.Save(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update item in cart", "success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item updated in cart successfully", "success": true, "data": cartItem})
}

// GetCart retrieves the cart for the user
func GetCart(c *gin.Context) {
	log.Println("Request received for retrieving the cart")
	userID := c.GetUint("user_id")

	var cartItems []models.Cart
	if err := models.DB.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve cart", "success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart retrieved successfully", "success": true, "data": cartItems})
}
