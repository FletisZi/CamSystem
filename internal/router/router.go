package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/pprof"
)

func Initialize() {
	router := gin.Default()

	pprof.Register(router)

	InitializeRoutes(router)

	router.Run(":8080")
}
