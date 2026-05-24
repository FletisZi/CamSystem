package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"camsystem/internal/models"
	"camsystem/internal/schemas"

	"github.com/gin-gonic/gin"
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
	fmt.Printf("Recebido: %+v\n", cam)

	// salva no banco
	if err := models.SaveCameraToDB(&cam); err != nil {
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

func DeleteCamera(c *gin.Context) {

	var id int
	idStr := c.Param("id")

	// converte string -> int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID inválido",
			"details": err.Error(),
		})
		return
	}

	// var cam schemas.CameraRequest

	// if err := c.ShouldBindJSON(&cam); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error":   "JSON inválido",
	// 		"details": err.Error(),
	// 	})
	// 	return
	// }

	// salva no banco
	if err := models.DeleteCameraFromDB(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "erro ao deletar câmera",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "câmera deletada com sucesso",
	})
}

func ListCameras(c *gin.Context) {

	cameras, err := models.GetAllCamerasFromDB()

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
