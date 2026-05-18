package db

import (
	"log"

	bolt "go.etcd.io/bbolt"
	"fmt"
	"os"
	
)

var DB *bolt.DB

func InitDB() {
	var err error

	wd, _ := os.Getwd()

	fmt.Println("===================================")
	fmt.Println("WORKDIR:", wd)
	fmt.Println("DATABASE:", wd + "/cameras.db")
	fmt.Println("===================================")

	DB, err = bolt.Open("cameras.db", 0600, nil)
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