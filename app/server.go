package main

import (
	"log"

	"github.com/YutoMizutani/gohome/app/application/util"
	"github.com/YutoMizutani/gohome/app/infrastructure"
)

func main() {
	util.PrepareEnv()
	err := infrastructure.Router.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
