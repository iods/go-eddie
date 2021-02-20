package env

import (
	"os"

	"github.com/iods/go-eddie/internal/errors"
	"github.com/joho/godotenv"
)

func Get(key string) string {
	err := godotenv.Load("../../.env")
	errors.Handle("error loading the .env file", err)
	return os.Getenv(key)
}