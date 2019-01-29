package common

import (
	"fmt"

	"net/http"
)

const (
	templateDir = "assets/template/common/"
)

// this is HomePage function serve index.html file from template dir
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, templateDir + "index.html")
}

func commonHome(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello from common index")
}
