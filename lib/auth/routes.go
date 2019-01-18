package auth

import (
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/register", Register).Methods("GET").Name("auth:register")
	r.HandleFunc("/login", Login).Methods("GET").Name("auth:login")
}
