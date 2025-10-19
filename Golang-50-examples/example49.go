// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Example: HTTP client
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response headers:")
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}
