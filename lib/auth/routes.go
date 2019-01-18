package auth

import (
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/register", Register).Methods("GET", "POST").Name("auth:register")
	r.HandleFunc("/login", Login).Methods("GET", "POST").Name("auth:login")
	r.HandleFunc("/profile", Profile).Methods("GET").Name("auth:profile")
}
