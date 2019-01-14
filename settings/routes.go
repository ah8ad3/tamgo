package settings

import "github.com/gorilla/mux"
import "../apps/common"

func routes() (router mux.Route){
	r := mux.NewRouter()

	r.HandleFunc("/", homePage).Methods("GET")

	return r
}
