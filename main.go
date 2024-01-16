package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    // Create a new HTTP request multiplexer
    mux := http.NewServeMux()

    // Register a handler function for the root path and wrap it with the logMetrics middleware
    mux.HandleFunc("/", logMetrics(func(w http.ResponseWriter, r *http.Request) {
        // Send a response to the client
        fmt.Fprint(w, "Hello, World! This is a cool web test!")
    }))

    // Create a new HTTP server with specified timeout settings
    server := &http.Server{
        Addr:         ":8080",
        Handler:      mux,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  15 * time.Second,
    }

    // Start the server in a new goroutine
    go func() {
        fmt.Println("Server is running on http://localhost:8080")
        // Listen and serve on the specified address and handle any errors
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("ListenAndServe(): %v", err)
        }
    }()

    // Set up graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    // Shutdown the server when a signal is received
    if err := server.Shutdown(nil); err != nil {
        log.Fatal("Server forced to shutdown:", err)
    }

    // Log server exit
    log.Println("Server exiting")
}

// responseRecorder is a custom http.ResponseWriter that records the status code and response size
type responseRecorder struct {
    http.ResponseWriter
    statusCode int
    size       int
}

// newResponseRecorder creates a new responseRecorder wrapping the given ResponseWriter
func newResponseRecorder(w http.ResponseWriter) *responseRecorder {
    return &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
}

// WriteHeader captures the status code for logging
func (rec *responseRecorder) WriteHeader(code int) {
    rec.statusCode = code
    rec.ResponseWriter.WriteHeader(code)
}

// Write captures the size of the response data for logging
func (rec *responseRecorder) Write(b []byte) (int, error) {
    size, err := rec.ResponseWriter.Write(b)
    rec.size += size
    return size, err
}

// logMetrics is a middleware that logs request details and metrics
func logMetrics(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Record the start time
        start := time.Now()
        // Create a new recorder to capture response details
        rec := newResponseRecorder(w)
        // Process the request
        next(rec, r)
        // Calculate the duration of the request
        duration := time.Since(start)

        // Log the request details including method, path, status code, response size, and duration
        log.Printf("Request: %s %s, Status: %d, Size: %d bytes, Duration: %s\n", r.Method, r.URL.Path, rec.statusCode, rec.size, duration)
    }
}
