package server

import (
	"fmt"
	"github.com/ah8ad3/tamgo/apps/common"
	"github.com/ah8ad3/tamgo/lib/auth"
	"github.com/ah8ad3/tamgo/settings"
	"log"
	"net/http"
)

const (
	// InfoColor blue
	InfoColor    = "\033[1;34m%s\033[0m"
	//NoticeColor  = "\033[1;36m%s\033[0m"
	//WarningColor = "\033[1;33m%s\033[0m"
	//ErrorColor   = "\033[1;31m%s\033[0m"
	//DebugColor   = "\033[0;36m%s\033[0m"
)

// Server is all settings and config functions
// also serve main http server
func Server() {
	// load all routes here
	r := settings.Routes()

	//load all settings here
	db := settings.Config()

	// need to set db to all models and apps here like this
	common.SetDB(db)
	auth.SetDB(db)


	// run server here
	fmt.Printf(InfoColor, "server run at localhost:8000")
	fmt.Println("")
	log.Fatal(http.ListenAndServe(":8000", r))
}
