package stream_manager

import (
	"sync"
	"camsystem/internal/camera"
)

type StreamManager struct {
	Cameras map[int]*camera.Camera
	Mu      sync.RWMutex
}

func NewStreamManager() *StreamManager {
	return &StreamManager{
		Cameras: make(map[int]*camera.Camera),
	}
}