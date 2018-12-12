package handler

import (
	"apptastic/dashboard/session"
	"net/http"

	"github.com/gorilla/sessions"
)

func LogoutHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		redirectTarget := "/login"

		session := session.Instance(r)
		session.Options.MaxAge = -1
		sessions.Save(r, w)

		http.Redirect(w, r, redirectTarget, 302)
	})
}
