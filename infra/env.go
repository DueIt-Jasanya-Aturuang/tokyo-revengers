package infra

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed load env | err : %v", err)
	}
}
