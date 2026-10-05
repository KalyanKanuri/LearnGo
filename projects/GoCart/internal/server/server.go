package server

import (
	"github.com/KalyanKanuri/GoCart/internal/config"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Server struct {
	cfg    *config.Config
	dbConn *gorm.DB
	logger *zerolog.Logger
}

func New(cfg *config.Config, dbConn *gorm.DB, logger *zerolog.Logger) *Server {
	return &Server{
		cfg:    cfg,
		dbConn: dbConn,
		logger: logger,
	}
}
