package main

import (
	"log"

	"github.com/YutoMizutani/gohome/app/infrastructure"

	"github.com/joho/godotenv"
)

func prepareEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	prepareEnv()
	err := infrastructure.Router.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
