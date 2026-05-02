package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := env("PORT", "8080")
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "iplusone Go site placeholder")
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	log.Printf("starting server on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
