package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func faviconHandler(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
