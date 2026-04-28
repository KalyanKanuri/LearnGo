package main

import (
	"learning/go/web/config"
	"learning/go/web/middleware"
	"learning/go/web/router"
	"net/http"
)

const PORT = ":8080"

func main() {
	app := config.GetAppConfig()

	mwManager := middleware.MWManager{
		App: app,
	}
	routeManager := router.RouteManager{
		App: app,
		MW:  &mwManager,
	}

	routeMux := routeManager.SetupRouter()
	mwStack := middleware.CreateMWStack(
		mwManager.LoggingMW,
		app.Session.Enable,
	)

	handler := mwStack(routeMux)

	app.Log.Info("Server Started on port %s", PORT)
	if err := http.ListenAndServe(PORT, handler); err != nil {
		app.Log.Critical("Error while starting the server %+v", err.Error())
	}
}
