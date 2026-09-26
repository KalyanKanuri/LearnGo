// Package db provides a function to create a new database connection using GORM and PostgreSQL.
package db

import (
	"fmt"

	"github.com/KalyanKanuri/GoCart/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// New creates a new GORM database connection using the provided database configuration.
func New(dbConfig *config.DatabaseConfig) (*gorm.DB, error) {
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

	return db, nil
}
