package main

import (
	"golang-gin-jwt/database"
	"golang-gin-jwt/router"
)

func main(){
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}