package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadENV(key string) string {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("failed to load environment")
	}
	dataSource := os.Getenv(key)
	return dataSource
}
