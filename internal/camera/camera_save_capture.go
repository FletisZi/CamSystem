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

	// cmd := exec.Command("ffmpeg",
	// 	"-f", "mpegts",
	// 	"-i", "pipe:0",
	// 	"-c", "copy",
	// 	"-movflags", "+faststart",
	// 	filename,
	// )

	// cmdFrame := exec.Command(
	// 	"ffmpeg",
	// 	"-f", "mpegts",
	// 	"-i", "pipe:0",
	// 	"-vf", "fps=1",
	// 	"-q:v", "2",
	// 	"c:\\Users\\rssantos\\Videos\\camsystem\\frames\\frame_%03d.jpeg",
	// )

	cmd := exec.Command(
    "ffmpeg",

    "-f", "mpegts",
    "-i", "pipe:0",

    // saída MP4
    "-map", "0:v",
    "-c:v", "copy",
    "-movflags", "+faststart",
    filename,

    // saída JPEG
    "-map", "0:v",
    "-vf", "fps=1",
    "-c:v", "mjpeg",
    "-q:v", "2",
    "c:\\Users\\rssantos\\Videos\\camsystem\\frames\\frame_%03d.jpeg",
)

	// stdinFrame, err := cmdFrame.StdinPipe()
	// if err != nil {
	// 	return "", err
	// }

	// if err := cmdFrame.Start(); err != nil {
	// 	return "", err
	// }

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()

		for _, chunk  := range c.RecordingBuffer {
			if _, err := stdin.Write(chunk); err != nil {
				fmt.Printf("Erro ao escrever frame para FFmpeg: %v\n", err)
				return
			}
		}
	}()

	// go func() {
	// 	defer stdinFrame.Close()

	// 	for _, chunk := range c.RecordingBuffer {
	// 		if _, err := stdinFrame.Write(chunk); err != nil {
	// 			fmt.Printf("Erro ao escrever frame para FFmpeg: %v\n", err)
	// 			return
	// 		}
	// 	}
	// }()

	videoErr := cmd.Wait()
	// frameErr := cmdFrame.Wait()

	if videoErr != nil {
		fmt.Printf("Erro ao salvar vídeo: %v\n", videoErr)
		return "", videoErr
	}

	// if frameErr != nil {
	// 	fmt.Printf("Erro ao salvar frames: %v\n", frameErr)
	// 	return "", frameErr
	// }	

	return filename, nil
}
