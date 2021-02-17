package env

import (
	"os"

	"github.com/iods/go-eddie/internal/utils"
	"github.com/joho/godotenv"
)

func EnvVar(key string) string {
	err := godotenv.Load("../../.env")
	utils.HandleError("Error loading the .env file.", err)
	return os.Getenv(key)
}