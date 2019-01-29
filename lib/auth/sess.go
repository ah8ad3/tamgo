package auth

import "github.com/gorilla/sessions"

// session pointer to set session store
var Store *sessions.CookieStore

// how Store pointer set
func SetStore(store *sessions.CookieStore) {
	Store = store
}
