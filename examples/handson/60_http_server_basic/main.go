package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/echo", echoHandler)

	fmt.Println("起動: http://localhost:8060")
	fmt.Println("  GET /hello")
	fmt.Println("  GET /echo?msg=...")
	log.Fatal(http.ListenAndServe(":8060", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Hello, HTTP!")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	if msg == "" {
		msg = "(empty)"
	}
	fmt.Fprintln(w, "echo:", msg)
}
