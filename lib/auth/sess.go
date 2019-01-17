package auth

import "github.com/gorilla/sessions"

var Store *sessions.CookieStore

func SetStore(store *sessions.CookieStore) {
	Store = store
}
