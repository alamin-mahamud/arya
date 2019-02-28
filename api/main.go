package main

import (
	"log"
	"net/http"

	"github.com/alamin-mahamud/arya/api/pkg"
)

func main() {
	config, err := pkg.LoadConfig("./config")
	if err != nil {
		panic("Could not load the configuration")
	}

	r := pkg.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(config.Server.Port, nil))
}
