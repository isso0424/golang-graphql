package server

import (
	"log"
	"net/http"
)

func Serve() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
