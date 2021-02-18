package env

import (
	"os"

	"github.com/iods/go-eddie/internal/helpers/error"
	"github.com/joho/godotenv"
)

func Get(key string) string {
	err := godotenv.Load("../../.env")
	error.Handle("error loading the .env file", err)
	return os.Getenv(key)
}