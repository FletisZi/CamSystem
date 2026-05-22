package main

import (
	"camsystem/internal/stream_manager"
	"camsystem/internal/capturador"
	"camsystem/internal/infra/db"
	"camsystem/internal/router"
	"camsystem/internal/monitor"
)

func main() {
	db.InitDB()
	

	manager := stream_manager.NewStreamManager()

	go capturador.Initialize(manager)

	monitor.StartMonitor()

	router.Initialize()
}