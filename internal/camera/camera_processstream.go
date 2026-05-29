package camera

import (
	"fmt"
	"io"
	"time"
)

func (c *Camera) ProcessStream(stdout io.ReadCloser) error {
	buf := make([]byte, 1024*64)
	const maxFrames = 30000000

	for {

		n, err := stdout.Read(buf)

		if err != nil {
			fmt.Printf("[Câmera %d] Erro na leitura do stream: %v\n", c.ID, err)
			return err // Se houver erro na leitura, sai do loop e avisa o "Vigia"
		}

		if n > 0 {
			c.Mu.Lock()
			c.IsRecording = true
			c.LastData = time.Now()
			c.Mu.Unlock()
		}

		frame := make([]byte, n)
		copy(frame, buf[:n])
		c.Buffer.Push(frame)

		if c.Recording {
			c.Mu.Lock()
			if len(c.RecordingBuffer) >= maxFrames {
				c.RecordingBuffer = c.RecordingBuffer[1:]
			}
			c.RecordingBuffer = append(c.RecordingBuffer, frame)
			c.Mu.Unlock()
			fmt.Printf("[Câmera %d] Gravando frame de %d bytes...\n", c.ID, n)
		} else {
			// fmt.Printf("[Câmera %d] Não está gravando %d bytes...\n", c.ID, n)
		}
	}
}
