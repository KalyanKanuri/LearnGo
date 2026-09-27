// Package db provides a function to create a new database connection using GORM and PostgreSQL.
package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/KalyanKanuri/GoCart/internal/config"
	"github.com/rs/zerolog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DataBase struct {
	db  *gorm.DB
	log *zerolog.Logger
}

// New creates a new GORM database connection using the provided database configuration.
func New(dbConfig *config.DatabaseConfig, log *zerolog.Logger) (*DataBase, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBUser,
		dbConfig.DBPassword,
		dbConfig.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	gcdb, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize DB: %w", err)
	}

	if err := gcdb.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB %w", err)
	}

	gcdb.SetMaxOpenConns(10)
	gcdb.SetMaxIdleConns(5)
	gcdb.SetConnMaxIdleTime(3 * time.Minute)
	gcdb.SetConnMaxLifetime(30 * time.Minute)

	return &DataBase{db: db, log: log}, nil
}

func (d *DataBase) GetDB() (*sql.DB, error) {
	return d.db.DB()
}
