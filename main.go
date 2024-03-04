package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	. "nuiteApi/controller"
)

func init() {
	godotenv.Load()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/hotels", GetHotels)
	return router
}

func main() {
	router := setupRouter()
	router.Run()
}
