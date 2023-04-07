package controllers

import (
	"fmt"
	"golang-gin-jwt/database"
	"golang-gin-jwt/helpers"
	"golang-gin-jwt/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID

	err := db.Debug().Create(&Product).Error
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Product)
}

func UpdateProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Product := models.Product{}
	productID, _ := strconv.Atoi(c.Param("productID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	
	Product.UserID = userID
	
	Product = models.Product{
		Title: Product.Title,
		Description: Product.Description,
	}

	err := db.Model(&Product).Where("id = ?", productID).Updates(Product).Error
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Product)
}

func GetProduct(c *gin.Context){
	productID, _ := strconv.Atoi(c.Param("productID"))
	Product := models.Product{}
	db := database.GetDB()

	err := db.First(&Product, productID).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Product with id %v not found", productID),
			})
		return
	}
	c.JSON(http.StatusOK, Product)	
}

func DeleteProduct(c *gin.Context){
	productID, _ := strconv.Atoi(c.Param("productID"))
	Product := models.Product{}
	db := database.GetDB()

	err := db.First(&Product, productID).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Product with id %v not found", productID),
			})
		return
	}

	if err :=db.Delete(&Product).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"Product deleted successfully",
	})
}