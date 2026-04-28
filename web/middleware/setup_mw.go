package middleware

import (
	"learning/go/web/config"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func CreateMWStack(mws ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(mws) - 1; i >= 0; i-- {
			mw := mws[i]
			next = mw(next)
		}
		return next
	}
}

type MWManager struct {
	App *config.AppConfig
}

func (mw MWManager) LoggingMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mw.App.Log.Debug("%s %s %s %s", r.Method, r.Proto, r.Host, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (mw MWManager) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authKey, err := r.Cookie("authKey")
		if err != nil {
			mw.App.Log.Error(err.Error())
		}

		switch authKey.Value {
		case "":
			mw.App.Log.Error("Unauthorized Access, empty cookie")
			http.Error(w, "Unauthorized, empty cookie", http.StatusUnauthorized)
			return
		case "1234":
			next.ServeHTTP(w, r)
		default:
			mw.App.Log.Debug("Unauthorized Access, incorrect cookie")
			http.Error(w, "Unauthorized access, cookie not set", http.StatusUnauthorized)
			return
		}
	})
}
