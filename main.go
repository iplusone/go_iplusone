package main

import (
	"log"
	"net/http"
	"os"

	"github.com/iplusone/go_iplusone/internal/site"
)

func main() {
	port := env("PORT", "8021")
	app := site.New()

	log.Printf("starting server on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, app.Handler()))
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
