package auth

import "github.com/gorilla/sessions"

// Store session pointer to set session store
var Store *sessions.CookieStore

// SetStore how Store pointer set
func SetStore(store *sessions.CookieStore) {
	Store = store
}
