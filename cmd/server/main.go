// main_server/main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	// Define the static directory
	fs := http.FileServer(http.Dir("./static"))

	// Serve static files
	http.Handle("/", fs)

	// Start the server on port 8080
	log.Println("Server is running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
