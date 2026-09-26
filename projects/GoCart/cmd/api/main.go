// Package main starts the GoCart API server.
package main

import (
	"github.com/KalyanKanuri/GoCart/internal/config"
	"github.com/KalyanKanuri/GoCart/internal/db"
	"github.com/KalyanKanuri/GoCart/internal/logger"
	"github.com/KalyanKanuri/GoCart/internal/server"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.New()
	log.Info().Msg("Starting GoCart API server...")

	conf, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	dbClient, err := db.New(&conf.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	dbConn, err := dbClient.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get database connection")
	}
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close database connection")
		}
	}()

	gin.SetMode(conf.App.GinMode)
	log.Info().Msgf("Started server on %s:%s", conf.App.AppHost, conf.App.AppPort)

	_ = server.New(conf, dbConn, &log)
}
