package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	//Store is the cookie store
	Store *sessions.CookieStore
	//Name is cookie name
	Name string
)

// SessionConfig stores session level information
type SessionConfig struct {
	Options   sessions.Options `json:"Options"`   // Pulled from: http://www.gorillatoolkit.org/pkg/sessions#Options
	Name      string           `json:"Name"`      // Name for: http://www.gorillatoolkit.org/pkg/sessions#CookieStore.Get
	SecretKey string           `json:"SecretKey"` // Key for: http://www.gorillatoolkit.org/pkg/sessions#CookieStore.New
}

// Configure the session cookie store
// global
func Configure(s SessionConfig) {
	Store = sessions.NewCookieStore([]byte(s.SecretKey))
	Store.Options = &s.Options
	Name = s.Name
}

// Instance returns a new session, never returns an error
func Instance(r *http.Request) *sessions.Session {
	session, _ := Store.Get(r, Name)
	return session
}

// Empty deletes all the current session values
func Empty(sess *sessions.Session) {
	// Clear out all stored values in the cookie
	for k := range sess.Values {
		delete(sess.Values, k)
	}
}
