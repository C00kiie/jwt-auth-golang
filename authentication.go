package main

import (
	"fmt"

	gosql "github.com/ilibs/gosql/v2"
)

func authenticate(username string, password string) (string, error) {
	config := getInstance()
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", config.DatabaseUsername, config.DatabasePassword, config.DatabaseName)
	configs := make(map[string]*gosql.Config)

	configs["default"] = &gosql.Config{
		Enable:  true,
		Driver:  "mysql",
		Dsn:     dsn,
		ShowSql: true,
	}
	gosql.Connect(configs)
	thisCredentials := authentication_credenetials{
		Username: username,
	}
	fmt.Println(thisCredentials)
	var query string = fmt.Sprintf("select id, username, password from goauth.%s where username = ?", config.Authentication_table)
	gosql.Get(&thisCredentials, query, username)

	fmt.Println(thisCredentials)
	if CheckPasswordHash(password, thisCredentials.Password) && thisCredentials.Username == username {
		token, _ := CreateToken(uint64(thisCredentials.Id))
		return token, nil
	}
	return "", fmt.Errorf("invalid credentials for user %s", username)
}

func register(username string, password string) (int64, error) {
	config := getInstance()
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", config.DatabaseUsername, config.DatabasePassword, config.DatabaseName)
	configs := make(map[string]*gosql.Config)

	configs["default"] = &gosql.Config{
		Enable:  true,
		Driver:  "mysql",
		Dsn:     dsn,
		ShowSql: true,
	}
	gosql.Connect(configs)
	Hashedpassword, _ := HashPassword(password)

	thisCredentials := make([]authentication_credenetials, 0)
	gosql.Select(&thisCredentials, "Select * from authentication where username = ?", username)
	fmt.Println(thisCredentials)
	if len(thisCredentials) != 0 {
		return -1, fmt.Errorf("the username %s is already taken", username)
	}

	gosql.Table("authentication").Create(map[string]interface{}{
		"username": username,
		"password": Hashedpassword,
	})

	return 0, nil
}
