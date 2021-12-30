package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	gosql "github.com/ilibs/gosql/v2"
)

// this is what our authentication_table looks like
type authentication_credenetials struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

func (u *authentication_credenetials) TableName() string {
	config := getInstance()
	return config.Authentication_table
}

func (u *authentication_credenetials) PK() string {
	return "id"
}

func authenticate(username string, password string) bool {
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
	gosql.Get(thisCredentials, "select username, password where username = ?", thisCredentials.Username)

	fmt.Println(thisCredentials)
	os.Exit(0)

	if CheckPasswordHash(password, thisCredentials.Password) && thisCredentials.Username == username {
		return true
	}
	return false
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
	var thisCredentials authentication_credenetials = authentication_credenetials{
		Id:       0,
		Username: username,
		Password: Hashedpassword,
	}

	gosql.Get(&thisCredentials, "select id, username from authentication where username  = ? ", username)

	if thisCredentials.Id != 0 {
		return -1, fmt.Errorf("the username %s is already taken", username)
	}

	id, err := gosql.Model(&thisCredentials).Create()
	if err != nil {
		panic(err)
	}

	return id, nil
}

func main() {
	fmt.Println(register("unique", "password"))

}
