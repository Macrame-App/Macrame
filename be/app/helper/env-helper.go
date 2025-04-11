package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func EnvGet(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		return ""
	}
	return os.Getenv("VITE_" + key)
}
