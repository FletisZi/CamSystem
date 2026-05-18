package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"camsystem/infra/db"
	"camsystem/schemas"

	"github.com/gin-gonic/gin"
	bolt "go.etcd.io/bbolt"
)

func CreateCamera(c *gin.Context) {

	// CORRETO: sem ponteiro
	var cam schemas.CameraRequest

	// bind do JSON
	if err := c.ShouldBindJSON(&cam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "JSON inválido",
			"details": err.Error(),
		})
		return
	}

	// salva no banco
	if err := SaveCameraToDB(&cam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "erro ao salvar câmera",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "câmera salva com sucesso",
	})
}

func SaveCameraToDB(cam *schemas.CameraRequest) error {

	return db.DB.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("cameras"))

		// valida bucket
		if b == nil {
			return fmt.Errorf("bucket cameras não encontrado")
		}

		// converte struct -> JSON
		data, err := json.Marshal(cam)
		if err != nil {
			return err
		}

		// chave do banco
		key := []byte(fmt.Sprintf("%d", cam.ID))

		// salva
		return b.Put(key, data)
	})
}

func ListCameras(c *gin.Context) {

	cameras, err := GetAllCamerasFromDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "erro ao listar câmeras",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cameras": cameras,
	})
}

func GetAllCamerasFromDB() (map[int]schemas.CameraRequest, error) {

	cameras := make(map[int]schemas.CameraRequest)

	err := db.DB.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("cameras"))

		// valida bucket
		if b == nil {
			return fmt.Errorf("bucket cameras não encontrado")
		}

		return b.ForEach(func(k, v []byte) error {

			// converte chave -> int
			id, err := strconv.Atoi(string(k))
			if err != nil {
				return err
			}

			// converte JSON -> struct
			var cam schemas.CameraRequest

			if err := json.Unmarshal(v, &cam); err != nil {
				return err
			}

			cameras[id] = cam

			return nil
		})
	})

	return cameras, err
}