package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// GetEnv func to get env value from key ---
func GetEnv(key string) string {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)

}
