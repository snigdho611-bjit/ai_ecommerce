package controllers

import (
	"ecommerce/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateProductInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
}

type UpdateProductInput struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Stock       *int     `json:"stock"`
}

func CreateProduct(c *gin.Context) {
	log.Println("Request received for creating a new product")
	var input CreateProductInput
	log.Println(input)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "success": false, "error": "The required fields were not formatted properly"})
		return
	}

	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	}

	if err := models.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create product", "success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "success": true, "data": product})
}

func GetAllProducts(c *gin.Context) {
	log.Println("Request received for retrieving all products")
	var products []models.Product
	if err := models.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve products", "success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Products retrieved successfully", "success": true, "data": products})
}

func UpdateProduct(c *gin.Context) {
	log.Println("Request received for updating a product")
	id := c.Param("id")
	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "success": false, "error": err.Error()})
		return
	}

	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found", "success": false, "error": err.Error()})
		return
	}

	// Update fields if they are provided in the request body
	updatedFields := make(map[string]interface{})
	if input.Name != nil {
		updatedFields["name"] = *input.Name
	}
	if input.Description != nil {
		updatedFields["description"] = *input.Description
	}
	if input.Price != nil {
		updatedFields["price"] = *input.Price
	}
	if input.Stock != nil {
		updatedFields["stock"] = *input.Stock
	}

	if len(updatedFields) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No valid fields provided for update", "success": false})
		return
	}

	if err := models.DB.Model(&product).Updates(updatedFields).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update product", "success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "success": true, "data": product})
}

func GetProductByID(c *gin.Context) {
	log.Println("Request received for retrieving a product by ID")
	id := c.Param("id")
	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found", "success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product retrieved successfully", "success": true, "data": product})
}

func DeleteProductByID(c *gin.Context) {
	log.Println("Request received for deleting a product by ID")
	id := c.Param("id")
	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found", "success": false, "error": err.Error()})
		return
	}

	if err := models.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete product", "success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully", "success": true})
}

func GetFilteredProducts(c *gin.Context) {
	log.Println("Request received for retrieving filtered products")
	var products []models.Product

	// Get query parameters
	name := c.Query("name")
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max_price")
	minStock := c.Query("min_stock")
	maxStock := c.Query("max_stock")

	query := models.DB

	// Apply filters
	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	if minPrice != "" {
		minPriceFloat, err := strconv.ParseFloat(minPrice, 64)
		if err == nil {
			query = query.Where("price >= ?", minPriceFloat)
		}
	}
	if maxPrice != "" {
		maxPriceFloat, err := strconv.ParseFloat(maxPrice, 64)
		if err == nil {
			query = query.Where("price <= ?", maxPriceFloat)
		}
	}
	if minStock != "" {
		minStockInt, err := strconv.Atoi(minStock)
		if err == nil {
			query = query.Where("stock >= ?", minStockInt)
		}
	}
	if maxStock != "" {
		maxStockInt, err := strconv.Atoi(maxStock)
		if err == nil {
			query = query.Where("stock <= ?", maxStockInt)
		}
	}

	if err := query.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve products", "success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Filtered products retrieved successfully", "success": true, "data": products})
}
