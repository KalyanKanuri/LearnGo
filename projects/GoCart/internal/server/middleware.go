package server

import (
	"errors"
	"net/http"
	"strings"

	"github.com/KalyanKanuri/GoCart/internal/auth"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}

func (srv *Server) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHdr := ctx.GetHeader("Authorization")
		if authHdr == "" {
			UnauthorizedError(
				ctx,
				"Invalid Authentication Header",
				errors.New("invalid authorization header"),
			)
			ctx.Abort()
			return
		}
		tokenParts := strings.Split(authHdr, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			UnauthorizedError(
				ctx,
				"Invalid Authorization Header",
				errors.New("invalid authorization header"),
			)
			ctx.Abort()
			return
		}

		claims, err := auth.ValidateToken(tokenParts[1], srv.cfg.JWT.JWTSecret, "Access")
		if err != nil {
			UnauthorizedError(ctx, "Invalid token", err)
			ctx.Abort()
			return
		}
		ctx.Set("user_id", claims.UserID)
		ctx.Set("user_email", claims.Email)
		ctx.Set("user_role", claims.Role)

		ctx.Next()
	}
}
