package main

import (
	"log"
	"net/http"

	"github.com/alamin-mahamud/arya/api/pkg"
	"github.com/alamin-mahamud/arya/api/pkg/config"
	"github.com/alamin-mahamud/arya/api/pkg/database"
)

func main() {
	config, err := config.LoadConfig("./config")
	if err != nil {
		panic("Could not load the configuration")
	}

	_, err = database.New(config)

	r := pkg.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(config.Server.Port, nil))
}
