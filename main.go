package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World! This is a basic web server!")
}

func main() {
    http.HandleFunc("/", handler) // Set the handler function for the root path
    fmt.Println("Server is running on http://localhost:8080") // Print a message indicating where the server is running
    http.ListenAndServe(":8080", nil) // Start the server on port 8080
}
