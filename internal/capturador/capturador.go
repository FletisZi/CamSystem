package capturador

import (
	"camsystem/internal/models"
	"camsystem/internal/stream_manager"
	"camsystem/internal/tools"
	"fmt"
)

func Initialize(manager *stream_manager.StreamManager) {
	cameras := models.GetCameras()

	for _, cam := range cameras {
		if cam.IsActive != false {
			url := tools.MontarURL(cam)
			fmt.Printf("Adicionando câmera ID %d com URL: %s\n", cam.ID, url)
			manager.AddCamera(int(cam.ID), cam.Name, url)
		}

	}

	// data, err := models.GetAllCamerasFromDB()

	// if err != nil {
	// 	panic("Erro ao acessar banco de dados: " + err.Error())
	// }

	// for _, cam := range data {
	// 	if cam.Active != nil && *cam.Active {
	// 		manager.AddCamera(cam.ID, cam.URL)
	// 	}
	// }

	// O programa principal fica rodando...
	select {}
}
