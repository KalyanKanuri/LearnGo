package session

import (
	"net/http"
	"time"

	"github.com/golangcollege/sessions"
)

var session *sessions.Session

func NewSession() *sessions.Session {
	session = sessions.New([]byte("kjaSFVUHdgvbdsmnvZXhgvcASN"))
	session.Lifetime = 1 * time.Hour
	return session
}

func GetFromSession(r *http.Request, session *sessions.Session, key string) any {
	val := session.Get(r, key)
	return val
}

func PutInSession(r *http.Request, session *sessions.Session, key_val map[string]any) {
	for key, val := range key_val {
		session.Put(r, key, val)
	}
}

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "authKey",
		Value: "1234",
	}
	http.SetCookie(w, cookie)
}
