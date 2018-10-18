package util

import (
	"os"

	"github.com/joho/godotenv"
)

func PrepareEnv(filenames ...string) error {
	err := godotenv.Load(filenames...)
	return err
}

func Getenv(key string) string {
	value := os.Getenv(key)
	return value
}
