package camera

import (
	"camsystem/internal/infra/ffmpeg"
	"fmt"
	"time"
)

func (c *Camera) StartCapture() {

	// go c.MonitorStream()

	go func() {
		for {

			fmt.Printf("[Câmera %d] Tentando conectar em: %s\n", c.ID, c.URL)

			cmd, stdout, err := ffmpeg.StartFFmpeg(c.URL)

			if err != nil {
				fmt.Printf("[Câmera %d] erro ao iniciar ffmpeg: %v\n", c.ID, err)
				time.Sleep(5 * time.Second)
				continue
			}

			err = c.ProcessStream(stdout)

			fmt.Printf("[Câmera %d] stream encerrado: %v\n", c.ID, err)

			// c.Mu.Lock()
			// c.IsRecording = false
			// c.Mu.Unlock()

			cmd.Wait()

			time.Sleep(5 * time.Second)
		}
	}()
}
