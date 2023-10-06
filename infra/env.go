package infra

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed load env | err : %v", err)
	}

	AccountUrl = os.Getenv("ACCOUNT_URL")
}

var (
	AccountUrl string
)
