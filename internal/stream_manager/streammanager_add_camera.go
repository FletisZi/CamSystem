package stream_manager

import (
	"fmt"
	"camsystem/internal/camera"
)

func (m *StreamManager) AddCamera(id int, name string, url string) {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	newCam := camera.NewCamera(id,name, url)
	m.Cameras[id] = newCam

	go newCam.StartCapture()

	fmt.Printf("[Manager] Câmera %d adicionada e iniciando...\n", id)
}