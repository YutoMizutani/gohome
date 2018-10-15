package main

import (
	"log"

	"github.com/YutoMizutani/gohome/app/infrastructure"
)

func main() {
	err := infrastructure.Router.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
