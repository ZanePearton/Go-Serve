package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Create a new ServeMux (router)
	mux := http.NewServeMux()

	// Register a handler function for the root URL
	mux.HandleFunc("/", rootHandler)

	// Create a new HTTP server
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: logRequest(mux), // Wrap the handler with the logging middleware
	}

	// Start the server and log if there's an error
	fmt.Println("Server is running on http://127.0.0.1:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

// rootHandler responds to requests at the root URL
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World! This is a cool web test!")
}

// logRequest is a middleware that logs each request
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("Request: %s %s, Duration: %s\n", r.Method, r.URL.Path, duration)
	})
}
