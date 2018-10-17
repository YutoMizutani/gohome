package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func PrepareEnv(filenames ...string) {
	err := godotenv.Load(filenames...)

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Getenv(key string) (value string) {
	value = os.Getenv(key)
	return
}
