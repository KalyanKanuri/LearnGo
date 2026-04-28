package router

import (
	"learning/go/web/config"
	"learning/go/web/handlers"
	"learning/go/web/middleware"
	"learning/go/web/ui"
	"net/http"
)

type RouteManager struct {
	App *config.AppConfig
	MW  *middleware.MWManager
}

func (rm RouteManager) SetupRouter() *http.ServeMux {
	handlerManager := handlers.HandlerManager{
		App: rm.App,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handlerManager.Greet(ui.NewTemplate()))
	mux.Handle("/about", rm.MW.AuthMiddleware(handlerManager.About(ui.NewTemplate())))
	return mux
}
