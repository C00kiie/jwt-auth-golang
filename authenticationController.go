package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func authenticate_wrapper(c *gin.Context) {
	var u authentication_credenetials

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:
	token, _ := authenticate(u.Username, u.Password)
	if token == "" {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	c.JSON(http.StatusOK, token)
}

func register_wrapper(c *gin.Context) {
	var u authentication_credenetials

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	fmt.Println(u)
	//compare the user from the request, with the one we defined:
	_, err := register(u.Username, u.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "this username is already taken")
		return
	}

	c.JSON(http.StatusOK, "Successfully registered")
}
