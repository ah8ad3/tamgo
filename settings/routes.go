package settings

import "github.com/gorilla/mux"
import "github.com/ah8ad3/tamgo/apps/common"

func Routes() (router *mux.Router){
	r := mux.NewRouter()

	r.HandleFunc("/", common.HomePage).Methods("GET")

	return r
}
