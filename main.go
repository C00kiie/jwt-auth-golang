package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", authenticate_wrapper)
	router.POST("/register", register_wrapper)
	router.Run()
}
