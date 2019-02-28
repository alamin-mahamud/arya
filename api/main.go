package main

import (
	"net/http"
	"strconv"

	"github.com/alamin-mahamud/arya/api/pkg"
)

const port = "8080"

func main() {
	config, err := pkg.LoadConfig("./config")
	if err != nil {
		panic("Could not load the configuration")
	}

	r := pkg.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":"+strconv.Itoa(config.ServerPort), nil)
}
