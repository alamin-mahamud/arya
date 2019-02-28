package main

import (
	"net/http"

	"github.com/alamin-mahamud/arya/api/pkg"
)

const port = "8080"

func main() {
	r := pkg.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
