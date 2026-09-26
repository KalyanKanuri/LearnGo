package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App    AppConfig
	DB     DatabaseConfig
	Upload UploadConfig
	AWS    AWSConfig
	JWT    JWTConfig
}

type AppConfig struct {
	AppHost string
	AppPort string
	GinMode string
}

type UploadConfig struct {
	MaxUploadSize int64
	UploadPath    string
}

type DatabaseConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

type AWSConfig struct {
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	AWSRegion          string
	AWSS3Bucket        string
	AWSS3Endpoint      string
}

type JWTConfig struct {
	JWTSecret         string
	JWTExpiration     time.Duration
	RefreshExpiration time.Duration
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	jwtExpiration, err := time.ParseDuration(requiredEnv("JWT_EXPIRATION"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXPIRATION: %w", err)
	}

	refreshExpiration, err := time.ParseDuration(requiredEnv("REFRESH_EXPIRATION"))
	if err != nil {
		return nil, fmt.Errorf("invalid REFRESH_EXPIRATION: %w", err)
	}

	maxUploadSize, err := strconv.ParseInt(optionalEnv("MAX_UPLOAD_SIZE", "10485760"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid MAX_UPLOAD_SIZE: %w", err)
	}

	cfg := &Config{
		App: AppConfig{
			AppHost: optionalEnv("APP_HOST", "localhost"),
			AppPort: optionalEnv("APP_PORT", "8080"),
			GinMode: optionalEnv("GIN_MODE", "debug"),
		},
		DB: DatabaseConfig{
			DBHost:     requiredEnv("DB_HOST"),
			DBPort:     optionalEnv("DB_PORT", "5432"),
			DBUser:     requiredEnv("DB_USER"),
			DBPassword: requiredEnv("DB_PASSWORD"),
			DBName:     requiredEnv("DB_NAME"),
		},
		AWS: AWSConfig{
			AWSAccessKeyID:     requiredEnv("AWS_ACCESS_KEY_ID"),
			AWSSecretAccessKey: requiredEnv("AWS_SECRET_ACCESS_KEY"),
			AWSRegion:          optionalEnv("AWS_REGION", "us-east-1"),
			AWSS3Bucket:        requiredEnv("AWS_S3_BUCKET"),
			AWSS3Endpoint:      optionalEnv("AWS_S3_ENDPOINT", "http://localhost:9000"),
		},
		JWT: JWTConfig{
			JWTSecret:         requiredEnv("JWT_SECRET"),
			JWTExpiration:     jwtExpiration,
			RefreshExpiration: refreshExpiration,
		},
		Upload: UploadConfig{
			MaxUploadSize: maxUploadSize,
			UploadPath:    optionalEnv("UPLOAD_PATH", "./uploads"),
		},
	}

	return cfg, nil
}

func requiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("missing required env: %s", key))
	}
	return value
}

func optionalEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
