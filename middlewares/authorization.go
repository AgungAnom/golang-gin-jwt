package middlewares

import (
	"golang-gin-jwt/database"
	"golang-gin-jwt/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error" : "Bad Request",
				"message" : "Invalid Parameter",
			})
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		userRole := string(userData["role"].(string))
		Product := models.Product{}

		err = db.Select("user_id").First(&Product,uint(productId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Data Not Found",
				"message" : "Data doesn't exist",
			})
			return
		}

		if Product.UserID != userID && userRole != "admin"{
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Unauthorized",
				"message" : "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}