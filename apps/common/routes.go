package common

import (
	"github.com/gorilla/mux"
)

// all routes of this app to add in subRouter in main router
func Routes(r *mux.Router) {
	r.HandleFunc("/", commonHome).Methods("GET").Name("common:index")
}
