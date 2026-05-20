package capturador

import (
	"camsystem/internal/stream_manager"
	"camsystem/internal/models"
)

func Initialize(manager *stream_manager.StreamManager) {
	data, err := models.GetAllCamerasFromDB()

	if err != nil {
		panic("Erro ao acessar banco de dados: " + err.Error())
	}

	for _, cam := range data {
		manager.AddCamera(cam.ID, cam.URL)
	}

	// O programa principal fica rodando...
	select {}
}