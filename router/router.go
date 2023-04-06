package router

import (
	"golang-gin-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
	}

	return r
} 