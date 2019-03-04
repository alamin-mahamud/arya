package router

import (
	"net/http"

	"github.com/alamin-mahamud/arya/auth/pkg/handler"
	"github.com/gorilla/mux"
)

// New - ...
func New() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handler.Ping)
	r.HandleFunc("/ping-error", handler.PingError)
	return r
}
