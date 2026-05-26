package models

import (
	"camsystem/internal/infra/db"
	"camsystem/internal/schemas"
	"fmt"
)

func GetCameras() []schemas.Cameras {
	db := db.GetDB()

	var camera []schemas.Cameras
	if err := db.Find(&camera).Error; err != nil {
		fmt.Printf("Erro ao buscar câmeras: %v\n", err)
		return nil
	}
	return camera
}
