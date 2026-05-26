package handlers

import (
	"camsystem/internal/infra/db"
	"camsystem/internal/schemas"
	"camsystem/internal/tools"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCameras(c *gin.Context) {
	db := db.GetDB()

	var input schemas.Cameras

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "JSON inválido",
			"details": err.Error(),
		})
		return
	}
	if input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "O campo 'nome' é obrigatório",
		})
		return
	}

	fmt.Printf("Recebido input: %+v\n", input)

	input.PasswordEncrypted, _ = tools.Encrypt(input.PasswordEncrypted)

	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao criar câmera",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": input,
	})
}
