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

	router.GET("/", rootHandler)
	// fix to suppress unnecessary 404 logs
	router.GET("/favicon.ico", faviconHandler)

	InitAuthHandlers(srv)
	InitUserHandlers(srv)
	InitProductHandlers(srv)

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

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(srv))
		{
			users := protected.Group("/users")
			{
				users.GET("/profile", GetProfileHandler)
				users.PUT("/profile", UpdateProfileHandler)
			}
			categories := protected.Group("/categories")
			{
				categories.GET("", GetCategoriesHandler)
				categories.POST("", CreateCategoryHandler)
				categories.PUT("", UpdateCategoryHandler)
				categories.DELETE("", DeleteCategoryHandler)
			}
			products := protected.Group("/products")
			{
				products.GET("", func(ctx *gin.Context) {
					if _, exists := ctx.GetQuery("prod_id"); exists {
						GetProductHandler(ctx)
						return
					}
					GetProductsHandler(ctx)
				})
				products.POST("", AddProductHandler)
				products.PUT("", UpdateProductHandler)
				products.DELETE("", DeleteProductHandler)
			}
		}
	}

	return router
}
