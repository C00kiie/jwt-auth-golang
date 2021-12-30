package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

const filename string = ".env"

type config struct {
	Secret               string
	DatabaseName         string
	DatabaseUsername     string
	DatabasePassword     string
	DatabaseHost         string
	Authentication_table string
}

var lock = &sync.Mutex{}

var singleInstance *config

func load_configuration(filename string) *config {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	config := config{
		Secret:               os.Getenv("Secret"),
		DatabaseName:         os.Getenv("DatabaseName"),
		DatabaseUsername:     os.Getenv("DatabaseUsername"),
		DatabasePassword:     os.Getenv("DatabasePassword"),
		Authentication_table: os.Getenv("Authentication_table"),
	}

	return &config
}
func getInstance() *config {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = load_configuration(filename)
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
