package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func GenerateVideoFilePath(plate string) (string, error) {
	basePath := os.Getenv("VIDEO_STORAGE_PATH")

	if basePath == "" {
		return "", fmt.Errorf("VIDEO_STORAGE_PATH environment variable is not set")
	}

	plate = strings.ToUpper(plate)

	now := time.Now()
	dateFolder := now.Format("02-01-2006")
	timestamp := now.Format("2006-01-02_15-04-05")

	datePath := filepath.Join(basePath, dateFolder)
	platePath := filepath.Join(datePath, plate)

	if err := os.MkdirAll(platePath, os.ModePerm); err != nil {
		return "", err
	}

	filename := filepath.Join(platePath, fmt.Sprintf("%s.mp4", timestamp))
	return filename, nil
}
