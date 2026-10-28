package handlers

import (
	"net/http"

	"github.com/KalyanKanuri/GoCart/internal/server"
	"github.com/KalyanKanuri/GoCart/internal/server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthCheck struct {
	db *gorm.DB
}

func NewHC(srv *server.Server) *HealthCheck {
	return &HealthCheck{
		db: srv.DBConn,
	}
}

func (hc *HealthCheck) checkDBHealth() (bool, error) {
	dbCheck, err := hc.db.DB()
	if err != nil {
		return false, err
	}

	if err := dbCheck.Ping(); err != nil {
		return false, err
	}

	return true, nil
}

var healthCheck *HealthCheck

func InitHealthCheck(srv *server.Server) {
	healthCheck = NewHC(srv)
}

func rootHandler(ctx *gin.Context) {
	dbHealth, err := healthCheck.checkDBHealth()
	if err != nil {
		utils.InternalServerError(
			ctx,
			"Server Unhealthy",
			err,
		)
		return
	}

	dbStatus := "not ok"
	if dbHealth {
		dbStatus = "ok"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"db_status": dbStatus,
	},
	)
}

func faviconHandler(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
