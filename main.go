package main

import (
	"log"
	"net/http"
)

// Adding:
// - server
// - address
// - handler
// - calling ListenAndServe

func main() {
	mux := http.NewServeMux()

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
}
