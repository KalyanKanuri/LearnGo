// Package main starts the GoCart API server.
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KalyanKanuri/GoCart/internal/config"
	"github.com/KalyanKanuri/GoCart/internal/database"
	"github.com/KalyanKanuri/GoCart/internal/logger"
	"github.com/KalyanKanuri/GoCart/internal/server"
	"github.com/KalyanKanuri/GoCart/internal/server/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.New()
	log.Info().Msg("Starting GoCart API server...")

	conf, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	dbClient, err := database.New(&conf.DB, &log)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database connection")
	}

	gin.SetMode(conf.App.GinMode)
	srv := server.New(conf, dbClient.GetDB(), &log)
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", conf.App.AppPort),
		Handler:      handlers.SetupRoutes(srv),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Info().Msgf("Started server on %s:%s", conf.App.AppHost, conf.App.AppPort)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Failed to Start http server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info().Msg("Shutting Down Server")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to shutdown server")
	}

	log.Info().Msg("Server Shutdown complete")
}
