package main

import (
	"log"
	"net/http"
	"time"

	"github.com/YutoMizutani/gohome/app/application/util"
	"github.com/YutoMizutani/gohome/app/infrastructure"
)

func main() {
	err := util.PrepareEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := infrastructure.Router

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = s.ListenAndServe()

	if err != nil {
		log.Fatal(err)
		return
	}
}
