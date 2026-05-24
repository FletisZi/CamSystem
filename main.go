package main

import (
	"camsystem/internal/capturador"
	"camsystem/internal/infra/db"
	"camsystem/internal/router"
	"camsystem/internal/stream_manager"
)

func main() {
	db.InitDB()

	manager := stream_manager.NewStreamManager()

	go capturador.Initialize(manager)

	// monitor.StartMonitor()

	router.Initialize()
}
