package common

import (
	"fmt"

	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	product := Product{Code: "12", Price: 12}

	_ = product
	// this is how we save a record to DB
	//DB.Create(&product)

	w.Header().Set("Content-Type", "application")
	_, _ = fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func commonHome(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello from common index")
}
