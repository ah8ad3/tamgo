package common

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application")
	_, _ = fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
