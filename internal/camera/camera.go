package camera

import (

	"camsystem/internal/tools"
	"os/exec"
	"sync"
	"time"
)

type Camera struct {
	ID              int
	Name			string
	URL             string
	Cmd             *exec.Cmd
	Buffer          *tools.RingBuffer
	Recording       bool
	IsRecording     bool
	isStopping      bool
	LastData        time.Time
	RecordingBuffer [][]byte
	Placa           string
	Mu              sync.RWMutex
}

func NewCamera(id int, name string, url string) *Camera {
	return &Camera{
		ID:              id,
		Name:            name,
		URL:             url,
		Buffer:          tools.NewRingBuffer(12, 20),
		Recording:       false,
		RecordingBuffer: make([][]byte, 0),
	}
}