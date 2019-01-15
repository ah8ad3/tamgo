package common

import (
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("", commonHome).Methods("GET").Name("common:index")
}
