package handlers

import (
	"github.com/KalyanKanuri/GoCart/internal/server"
	"github.com/KalyanKanuri/GoCart/internal/server/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(srv *server.Server) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CorsMiddleware())
	// router.Use(srv.AuthMiddleware())

	router.GET("/", rootHandler)
	// fix to suppress unnecessary 404 logs
	router.GET("/favicon.ico", faviconHandler)

	InitAuthHandlers(srv)
	api := router.Group("/api/v1")
	{
		// Register Auth Handlers
		auth := api.Group("/auth")
		{
			auth.POST("/register", registerHandler)
			auth.POST("/login", loginHandler)
			auth.POST("/refresh-token", refreshTokenHandler)
			auth.POST("/logout", logoutHandler)
		}
	}

	return router
}
