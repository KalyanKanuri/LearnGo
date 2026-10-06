package server

import (
	"github.com/KalyanKanuri/GoCart/internal/config"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Server struct {
	CFG    *config.Config
	DBConn *gorm.DB
	Logger *zerolog.Logger
}

func New(cfg *config.Config, dbConn *gorm.DB, logger *zerolog.Logger) *Server {
	return &Server{
		CFG:    cfg,
		DBConn: dbConn,
		Logger: logger,
	}
}
