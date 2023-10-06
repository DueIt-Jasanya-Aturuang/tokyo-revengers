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
	AuthUrl = os.Getenv("AUTH_URL")
	FinanceUrl = os.Getenv("FINANCE_URL")

	AccountKey = os.Getenv("ACCOUNT_KEY")
	AuthKey = os.Getenv("AUTH_KEY")
	FinanceKey = os.Getenv("FINANCE_KEY")
	Key = os.Getenv("KEY")
}

var (
	AccountUrl string
	AuthUrl    string
	FinanceUrl string
)

var (
	AccountKey string
	AuthKey    string
	FinanceKey string
	Key        string
)
