// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Example: Simple HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! You requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
