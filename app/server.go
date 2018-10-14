package main

import "github.com/YutoMizutani/gohome/app/infrastructure"

func main() {
	router := infrastructure.Router
	router.Run()
}
