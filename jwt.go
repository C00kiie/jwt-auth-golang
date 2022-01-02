package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	gosql "github.com/ilibs/gosql/v2"
)

func VerifyToken(token string) bool {
	config := getInstance()
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", config.DatabaseUsername, config.DatabasePassword, config.DatabaseName)
	configs := make(map[string]*gosql.Config)

	configs["default"] = &gosql.Config{
		Enable:  true,
		Driver:  "mysql",
		Dsn:     dsn,
		ShowSql: true,
	}
	gosql.Connect(configs)
	var instance sessions_table = sessions_table{}
	err := gosql.Get(&instance, "select * from goauth.jwt_sessions where jwt_token = ? ", token)
	if err != nil {
		panic(err)
	}
	if instance.Created_at.Add(time.Hour).Unix() > time.Now().Unix() {
		return true
	}
	return false
}

func CreateToken(userid uint64) (string, error) {
	var err error
	//Creating Access Token

	config := getInstance()
	os.Setenv("ACCESS_SECRET", config.Secret)

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", config.DatabaseUsername, config.DatabasePassword, config.DatabaseName)
	configs := make(map[string]*gosql.Config)

	configs["default"] = &gosql.Config{
		Enable:  true,
		Driver:  "mysql",
		Dsn:     dsn,
		ShowSql: true,
	}
	gosql.Connect(configs)

	//this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	gosql.Table("jwt_sessions").Create(map[string]interface{}{
		"userId":     int(userid),
		"created_at": time.Now(),
		"jwt_token":  token,
	})

	return token, nil
}
