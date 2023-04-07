package router

import (
	"golang-gin-jwt/controllers"
	"golang-gin-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}
	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/",controllers.CreateProduct)
		productRouter.PUT("/:productID",middlewares.ProductAuthorization(),controllers.UpdateProduct)
		productRouter.GET("/:productID",middlewares.ProductAuthorization(),controllers.UpdateProduct)
		productRouter.DELETE("/:productID",middlewares.ProductAuthorization(),controllers.UpdateProduct)

	}


	return r
} 