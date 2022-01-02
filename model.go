package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// this is what our authentication_table looks like
type authentication_credenetials struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type sessions_table struct {
	Id         int       `json:"id" db:"id"`
	UserId     int       `json:"userId" db:"userId"`
	Jwt_token  string    `json:"jwt_token" db:"jwt_token"`
	Created_at time.Time `json:"created_at" db:"created_at"`
}

func (u *authentication_credenetials) TableName() string {
	return "jwt_sessions"
}

func (u *authentication_credenetials) PK() string {
	return "id"
}
func (u *sessions_table) TableName() string {
	config := getInstance()
	return config.Authentication_table
}

func (u *sessions_table) PK() string {
	return "id"
}
