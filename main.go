package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	. "nuiteApi/controller"
)

func init() {
	godotenv.Load()
}

func main() {
	router := gin.Default()
	router.GET("/hotels", GetHotels)
	router.Run()
}
