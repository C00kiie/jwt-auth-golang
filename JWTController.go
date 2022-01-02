package main

import (
	sjwt "github.com/brianvoe/sjwt"
	"github.com/gin-gonic/gin"
)

func verifyJwt(c *gin.Context) bool {
	token := c.GetHeader("token")
	config := getInstance()
	secretKey := []byte(config.Secret)
	hasVerified := sjwt.Verify(token, secretKey)

	if !hasVerified {
		return false
	}

	// Parse jwt
	claims, _ := sjwt.Parse(token)

	// Validate will check(if set) Expiration At and Not Before At dates
	err := claims.Validate()
	if err != nil {
		return false
	}
	return true
}
