package server

import (
	"github.com/ah8ad3/tamgo/settings"
	"log"
	"net/http"
)

func Server() {
	r := settings.Routes()
	log.Fatal(http.ListenAndServe(":8000", r))
}
