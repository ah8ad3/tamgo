package settings

import (
	"github.com/ah8ad3/tamgo/lib/auth"
	"github.com/gorilla/mux"
)
import "github.com/ah8ad3/tamgo/apps/common"

func Routes() (router *mux.Router){
	r := mux.NewRouter()
	r.HandleFunc("/", common.HomePage).Methods("GET").Name("index")

	r.HandleFunc("/coo", auth.MyHandler)
	r.HandleFunc("/check", auth.SessionCheck)

	// this is how we create our routes for apps
	common.Routes(r.PathPrefix("/common").Subrouter())
	return r
}
