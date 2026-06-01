package stream_manager

import (
	"fmt"
)

func (m *StreamManager) Start(id int) error {
	m.Mu.RLock()
	defer m.Mu.RUnlock()

	camera := m.Cameras[id]

	if camera == nil {
		return fmt.Errorf("câmera %d não encontrada", id)
	}

	data := camera.Buffer.GetAll()
	camera.RecordingBuffer = append(camera.RecordingBuffer, data...)

	fmt.Println("[Manager] Iniciando Captura da Câmera:", camera.Name)
	camera.Recording = true

	return nil
}

func (m *StreamManager) Stop(id int) error {
	m.Mu.RLock()
	defer m.Mu.RUnlock()

	camera := m.Cameras[id]
	if camera == nil {
		return fmt.Errorf("câmera %d não encontrada", id)
	}
	camera.Recording = false

	fmt.Println("[Manager] Parando Captura da Câmera:", camera.Name)

	filename, err := camera.SaveRecording()

	if err != nil {
		fmt.Printf("[Manager] Erro ao salvar gravação da câmera %d: %v\n", camera.ID, err)
		return err
	}
	// salvar metadados do evento no banco de dados aqui, associando o caminho do vídeo e a placa
	fmt.Printf("[Manager] Gravação da câmera %d salva com sucesso em: %s\n", camera.ID, filename)


	// if err != nil {
	// 	fmt.Printf("[Handler] Erro ao salvar gravação da câmera %d: %v\n", camera.ID, err)
	// }

	// salvar metadados do evento no banco de dados aqui, associando o caminho do vídeo e a placa

	camera.RecordingBuffer = nil

	return nil
}
