package controllers

import (
	"errors"
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

	err := db.Debug().Create(&User).Error
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

func UserLogin (c *gin.Context){
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password := User.Password


	err := db.Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.AbortWithError(http.StatusBadRequest, errors.New("Invalid password"))
		return
	}

	token, err := helpers.GenerateToken(User.ID, User.Email)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}