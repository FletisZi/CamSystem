package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/pprof"
	"camsystem/internal/stream_manager"
)

func Initialize( streamManager *stream_manager.StreamManager) {
	router := gin.Default()

	pprof.Register(router)

	InitializeRoutes(router,streamManager)

	router.Run("localhost:8080")
}
