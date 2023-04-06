package controllers

import (
	"golang-gin-jwt/database"
	"golang-gin-jwt/helpers"
	"golang-gin-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User :=models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Create(&User).Error
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id" : User.ID,
		"email": User.Email,
		"full_name":User.FullName,
		"role":User.Role,
	})
}