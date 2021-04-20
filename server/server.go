package server

import (
	"log"
	"net/http"
)

func Serve() {
	http.HandleFunc("/graphql", graphqlRoute)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
