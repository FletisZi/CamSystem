package tools

import (
	"camsystem/internal/schemas"
	"fmt"
	"strconv"
)

func MontarURL(camera schemas.Cameras) string {
	password, err := Decrypt(camera.PasswordEncrypted)
	if err != nil {
		fmt.Printf("Erro ao descriptografar senha: %v\n", err)
		return err.Error()
	}

	URL := fmt.Sprintf("rtsp://%s:%s@%s:%s%s", camera.Username, password, camera.IP, strconv.Itoa(camera.Port), camera.RTSPUrl)

	return URL

}
