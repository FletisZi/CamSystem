package tools

import (
	"crypto/rand"
	"os"
	"path/filepath"
)

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func gerarCodigo(tamanho int) (string, error) {
	bytes := make([]byte, tamanho)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i := range bytes {
		bytes[i] = chars[int(bytes[i])%len(chars)]
	}

	return string(bytes), nil
}

func GerarCodigoUnico(baseDir string) (string, error) {
	for {
		codigo, err := gerarCodigo(6)
		if err != nil {
			return "", err
		}

		pasta := filepath.Join(baseDir, codigo)

		if _, err := os.Stat(pasta); os.IsNotExist(err) {
			return codigo, nil
		}
	}
}