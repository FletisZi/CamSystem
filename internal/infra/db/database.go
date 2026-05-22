package db

import (
	"log"
	"os"
	bolt "go.etcd.io/bbolt"
	"github.com/joho/godotenv"

)

var DB *bolt.DB

func InitDB() {
	var err error

	// carrega .env
	err = godotenv.Load("./internal/config/.env")
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	// pega variável
	dbPath := os.Getenv("DATABASE_PATH")

	// cria pasta
	err = os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// abre banco
	DB, err = bolt.Open(dbPath+"/app.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}


	// cria bucket se não existir
	err = DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("cameras"))
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
	
}