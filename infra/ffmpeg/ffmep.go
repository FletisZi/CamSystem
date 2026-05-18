package ffmpeg

import (
	"fmt"
	"io"
	"os/exec"
)

func NewFFmpegCommand(url string) *exec.Cmd {
	return exec.Command(
		"ffmpeg",
		"-rtsp_transport", "tcp",
		"-timeout", "10000000", // 500ms
		"-i", url,
		"-c", "copy",
		"-f", "mpegts",
		"pipe:1",
	)
}

func StartFFmpeg(url string) (*exec.Cmd, io.ReadCloser, error) {
	cmd := NewFFmpegCommand(url)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Erro ao iniciar FFmpeg: %v\n", err)
		return nil, nil, err
	}

	return cmd, stdout, nil
}