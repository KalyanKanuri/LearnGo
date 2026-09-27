package server

import (
	"github.com/gin-gonic/gin"
)

func (srv Server) SetupRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CorsMiddleware())

	router.GET("/", rootHandler)
	// fix to suppress unnecessary 404 logs
	router.GET("/favicon.ico", faviconHandler)

	return router
}
