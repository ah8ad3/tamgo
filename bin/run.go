package bin

import (
	"../settings"
	"log"
	"net/http"
)

func main() {
	r := routes()
	log.Fatal(http.ListenAndServe(":8000", r))
}
