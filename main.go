package main

import (
	"fmt"
	"log"
	"net/http"
)

// main runs the server on port 8080
func main() {
	fmt.Println("Starting Go Task Tracker on http://localhost:8080")

	// Serve static frontend files from /static folder
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Start server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}