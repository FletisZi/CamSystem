package camera

import (
	"camsystem/internal/tools"
	"fmt"
	"os/exec"
)

func (c *Camera) SaveRecording(placa string) error {
	c.Mu.RLock()
	defer c.Mu.RUnlock()

	fmt.Println("Placa é", placa)

	// type Payload struct {
	// 	Placa string `json:"placa"`
	// }

	// var data Payload

	// err := json.Unmarshal([]byte(placa), &data)
	// if err != nil {
	// 	fmt.Println("Erro ao converter:", err)
	// 	return err
	// }

	// filename, err := tools.GenerateVideoFilePath(data.Placa)
	// if err != nil {
	// 	fmt.Println("Erro ao gerar caminho do vídeo:", err)
	// 	return err
	// }

	filename, err := tools.GenerateVideoFilePath(placa)
	if err != nil {
		fmt.Println("Erro ao gerar caminho do vídeo:", err)
		return err
	}

	fmt.Println("Salvando em:", filename)

	cmd := exec.Command("ffmpeg",
		"-f", "mpegts",
		"-i", "pipe:0",
		"-c", "copy",
		"-movflags", "+faststart",
		filename,
	)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		for _, frame := range c.RecordingBuffer {
			stdin.Write(frame)
		}
	}()

	return cmd.Wait()
}
