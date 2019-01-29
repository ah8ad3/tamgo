package auth

import (
	"github.com/gorilla/mux"
)

// all routes for auth app
func Routes(r *mux.Router) {
	r.HandleFunc("/register", Register).Methods("GET", "POST").Name("auth:register")
	r.HandleFunc("/login", Login).Methods("GET", "POST").Name("auth:login")
	r.HandleFunc("/logout", Logout).Methods("GET").Name("auth:logout")
	r.HandleFunc("/profile", Profile).Methods("GET").Name("auth:profile")
	r.HandleFunc("/jwt/check", CheckJwt).Methods("POST")
	r.HandleFunc("/jwt/sign", SignJWT).Methods("POST")
}
