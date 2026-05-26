package db

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("./internal/config/.env")
	if err != nil {
		log.Println("⚠ Arquivo .env não encontrado, usando variáveis do sistema")
	}
}
