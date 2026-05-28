package stream_manager

import (
	"fmt"
	
)

func (m *StreamManager) Start(id int) error {
	
	camera := m.Cameras[id]

	if camera == nil {
		return fmt.Errorf("câmera %d não encontrada", id)
	}

	fmt.Println("[Manager] Iniciando Captura da Câmera:", camera.Name)
	
	return nil
}


func (m *StreamManager) Stop(id int) {
	
	camera := m.Cameras[id]
	fmt.Println("[Manager] Parando Captura da Câmera:", camera.Name)
}