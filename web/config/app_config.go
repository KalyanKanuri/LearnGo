package config

import (
	"learning/go/web/logger"
	"learning/go/web/session"

	"github.com/golangcollege/sessions"
)

type AppConfig struct {
	Log     *logger.LogConfig
	Session *sessions.Session
}

func GetAppConfig() *AppConfig {
	return &AppConfig{
		Log:     logger.NewLogger(nil, logger.Debug),
		Session: session.NewSession(),
	}
}
