//go:build !js && !wasm
// +build !js,!wasm

package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve static files from the current directory
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Add WASM MIME type
	http.HandleFunc("/main.wasm", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/wasm")
		http.ServeFile(w, r, "main.wasm")
	})

	log.Print("Listening on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}