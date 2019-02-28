package main

import (
	"fmt"
	"net/http"
)

const port = "8080"

func main() {
	http.HandleFunc("/", simple)
	http.ListenAndServe(":"+port, nil)
}

func simple(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
