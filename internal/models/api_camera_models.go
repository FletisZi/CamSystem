package models

import (
	"encoding/json"
	"fmt"

	"camsystem/internal/infra/db"
	"camsystem/internal/schemas"
	"strconv"

	bolt "go.etcd.io/bbolt"
)

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

func DeleteCameraFromDB(id int) error {

	return db.DB.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("cameras"))

		// valida bucket
		if b == nil {
			return fmt.Errorf("bucket cameras não encontrado")
		}

		// chave
		key := []byte(strconv.Itoa(id))

		// verifica se existe antes de deletar
		if b.Get(key) == nil {
			return fmt.Errorf("camera com ID %d não encontrada", id)
		}

		// remove do banco
		if err := b.Delete(key); err != nil {
			return fmt.Errorf("erro ao deletar camera: %w", err)
		}

		return nil
	})
}
