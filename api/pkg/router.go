package pkg

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter - ...
func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/ping", Ping)
	return r
}