package server

import (
	"database/sql"

	"github.com/KalyanKanuri/GoCart/internal/config"
	"github.com/rs/zerolog"
)

type Server struct {
	cfg    *config.Config
	dbConn *sql.DB
	logger *zerolog.Logger
}

func New(cfg *config.Config, dbConn *sql.DB, logger *zerolog.Logger) *Server {
	return &Server{
		cfg:    cfg,
		dbConn: dbConn,
		logger: logger,
	}
}
