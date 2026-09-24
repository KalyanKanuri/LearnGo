// Package config provides functionality to load and manage application configuration from environment variables.
package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all the configuration settings for the application.
type Config struct {
	App    AppConfig
	DB     DatabaseConfig
	Upload UploadConfig
	AWS    AWSConfig
	JWT    JWTConfig
}

// AppConfig holds the configuration for the application.
type AppConfig struct {
	AppHost string `env:"APP_HOST" `
	AppPort string `env:"APP_PORT"`
	GinMode string `env:"GIN_MODE"`
}

// UploadConfig holds the configuration for file uploads.
type UploadConfig struct {
	MaxUploadSize int64  `env:"MAX_UPLOAD_SIZE"`
	UploadPath    string `env:"UPLOAD_PATH"`
}

// DatabaseConfig holds the configuration for the database connection.
type DatabaseConfig struct {
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
}

// AWSConfig holds the configuration for AWS services.
type AWSConfig struct {
	AWSAccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	AWSRegion          string `env:"AWS_REGION"`
	AWSS3Bucket        string `env:"AWS_S3_BUCKET"`
	AWSS3Endpoint      string `env:"AWS_S3_ENDPOINT"`
}

// JWTConfig holds the configuration for JWT authentication.
type JWTConfig struct {
	JWTSecret         string        `env:"JWT_SECRET"`
	JWTExpiration     time.Duration `env:"JWT_EXPIRATION"`
	RefreshExpiration time.Duration `env:"REFRESH_EXPIRATION"`
}

// Load reads the configuration from environment variables and returns a Config struct.
func Load() (*Config, error) {
	_ = godotenv.Load()
	jwtExpiration, err := time.ParseDuration(getEnv("JWT_EXPIRATION", "24h"))
	if err != nil {
		return nil, err
	}

	refreshExpiration, err := time.ParseDuration(getEnv("REFRESH_EXPIRATION", "72h"))
	if err != nil {
		return nil, err
	}

	maxUploadSize, err := strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE", "100"), 10, 64)
	if err != nil {
		return nil, err
	}

	return &Config{
		App: AppConfig{
			AppHost: getEnv("APP_HOST", "localhost"),
			AppPort: getEnv("APP_PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		DB: DatabaseConfig{
			DBHost:     getEnv("DB_HOST", "localhost"),
			DBPort:     getEnv("DB_PORT", "5432"),
			DBUser:     getEnv("DB_USER", "your_db_user"),
			DBPassword: getEnv("DB_PASSWORD", "your_db_password"),
			DBName:     getEnv("DB_NAME", "your_db_name"),
		},
		AWS: AWSConfig{
			AWSAccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", "your_aws_access_key"),
			AWSSecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", "your_aws_secret_key"),
			AWSRegion:          getEnv("AWS_REGION", "us-east-1"),
			AWSS3Bucket:        getEnv("AWS_S3_BUCKET", "your_aws_s3_bucket"),
			AWSS3Endpoint:      getEnv("AWS_S3_ENDPOINT", "http://localhost:9000"),
		},
		JWT: JWTConfig{
			JWTSecret:         getEnv("JWT_SECRET", "your_jwt_secret_key"),
			JWTExpiration:     jwtExpiration,
			RefreshExpiration: refreshExpiration,
		},
		Upload: UploadConfig{
			MaxUploadSize: maxUploadSize,
			UploadPath:    getEnv("UPLOAD_PATH", "./uploads"),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
