package handlers

import (
	"learning/go/web/config"
	"learning/go/web/session"
	"learning/go/web/ui"
	"net/http"
)

type HandlerManager struct {
	App *config.AppConfig
}

func (hm HandlerManager) Greet(tmpl *ui.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hm.App.Log.Debug("greet func called")
		ses_data := make(map[string]any)
		ses_data["AuthKey"] = "lksjckjsabcj"
		session.PutInSession(r, hm.App.Session, ses_data)
		session.SetCookie(w, r)
		tmpl.RenderTemplate(w, "hello.html", nil)
	}
}

func (hm HandlerManager) About(tmpl *ui.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hm.App.Log.Debug("about func called")
		val := session.GetFromSession(r, hm.App.Session, "AuthKey")
		hm.App.Log.Debug("AuthKey: %s", val.(string))
		tmpl.RenderTemplate(w, "about.html", nil)
	}
}
