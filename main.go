package main

import (
	"camsystem/infra/db"
	"camsystem/router"
)

func main() {
	db.InitDB()
	router.Initialize()
}