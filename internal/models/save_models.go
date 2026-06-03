package models

import (
	"camsystem/internal/infra/db"
	"camsystem/internal/schemas"
	"fmt"
)

func SaveMetaData(playload schemas.VideoRecordings) error {
	db := db.GetDB()
	if err := db.Create(&playload).Error; err != nil {
		fmt.Printf("Erro ao criar gravação de vídeo: %v\n", err)
		return err
	}
	return nil
}
