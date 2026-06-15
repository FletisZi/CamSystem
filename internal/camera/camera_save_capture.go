package camera

import (
	"camsystem/internal/tools"
	"fmt"
	"os/exec"
	
)

func (c *Camera) SaveRecording() (string, error) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()

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

	filename, err := tools.GenerateVideoFilePath("ABC123") 
	if err != nil {
		fmt.Println("Erro ao gerar caminho do vídeo:", err)
		return "", err
	}

	fmt.Println("Salvando em:", filename)

	cmd := exec.Command("ffmpeg",
		"-f", "mpegts",
		"-i", "pipe:0",
		"-c", "copy",
		"-movflags", "+faststart",
		filename,
	)

	cmdFrame := exec.Command("ffmpeg",
		"-f", "mpegts",
		"-i", "pipe:0",
		"-vf", "fps=1/2", // Extrai 1 frame por segundo
		"-q:v", "2",    // Define a qualidade dos frames (1 é a melhor qualidade, 31 é a pior)
		"-f", "image2pipe",
		"-vcodec", "mjpeg",
		// "-pix_fmt", "yuv420p",não sei o que é isso 
		"pipe:1",
		`c:\\Users\\rssantos\\Videos\\camsystem\\frames\\frame_%d.jpeg`, // Salva os frames como JPEGs numerados
		// "C:\\Users\\rssantos\\Videos\\camsystem\\frames\\frame_111.jpg",
	)

	stdinFrame, err := cmdFrame.StdinPipe()
	if err != nil {
		return "", err
	}

	if err := cmdFrame.Start(); err != nil {
		return "", err
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()

		for _, frame := range c.RecordingBuffer {
			stdin.Write(frame)
		}
	}()

	go func() {
		defer stdinFrame.Close()

		for _, frame := range c.RecordingBuffer {
			stdinFrame.Write(frame)
		}
	}()

	return filename, cmd.Wait()
}
