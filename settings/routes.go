package settings

import (
	"fmt"
	"github.com/ah8ad3/tamgo/lib/auth"
	"github.com/gorilla/mux"
	"net/http"
)
import "github.com/ah8ad3/tamgo/apps/common"

func listRoutes(r *mux.Router) {
	_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})
}

// this is main routes function that connect to all apps
func Routes() (router *mux.Router){
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", common.HomePage).Methods("GET").Name("index")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("assets/static/"))))

	// this is how we create our routes for apps
	common.Routes(r.PathPrefix("/common").Subrouter())
	auth.Routes(r.PathPrefix("/auth").Subrouter())

	// function for list all routes
	//listRoutes(r)

	return r
}
