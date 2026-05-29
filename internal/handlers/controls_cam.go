package handlers

import (
	"camsystem/internal/stream_manager"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StartCamera(manager *stream_manager.StreamManager) gin.HandlerFunc {

	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "ID inválido"})
			return
		}

		error := manager.Start(id)

		if error != nil {
			c.JSON(404, gin.H{
				"error": error.Error(),
			})
			return
		}

		c.JSON(200, gin.H{"message": "Câmera iniciada com sucesso"})
	}
}

func StopCamera(manager *stream_manager.StreamManager) gin.HandlerFunc {

	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "ID inválido"})
			return
		}

		error := manager.Stop(id)

		if error != nil {
			c.JSON(404, gin.H{
				"error": error.Error(),
			})
			return
		}

		c.JSON(200, gin.H{"message": "Câmera parada com sucesso"})

	}
}
