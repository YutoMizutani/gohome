package main

import (
	"log"

	"github.com/YutoMizutani/gohome/app/application/util"
	"github.com/YutoMizutani/gohome/app/infrastructure"
)

func main() {
	err := util.PrepareEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = infrastructure.Router.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
