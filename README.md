# Go-Serve

This simple web server, written in Go, responds with "Hello, World! This is a basic web server!" to all requests.

## Description

Go-Serve is a foundational project designed to provide an understanding of the basics of web serving in Go. It listens on port 8080 and serves a straightforward greeting message, perfect for beginners looking to get their feet wet in web development with Go.

## Getting Started

### Dependencies

- Go programming language

### Installing

Clone the repository to your local machine:

```bash
git clone https://github.com/ZanePearton/go-serve.git
```
## Running the Server

Navigate to the directory where you cloned the repo.
Execute the program:
```bash
go run server.go
```

### Go-Serve Server
1. **Package Declaration**: `package main` - This declares that this file is part of the main package, which is the executable package for a Go application.

2. **Imports**:
    - `fmt`: Importing the "fmt" package allows you to use functions for formatting text, including printing to the console.
    - `net/http`: This package provides HTTP client and server implementations, and you're using it to create the web server and handle HTTP requests.

3. **Handler Function**:
    - `func handler(w http.ResponseWriter, r *http.Request)`: This function is a handler that gets called every time you make an HTTP request to the server. It takes two parameters:
        - `w` (an `http.ResponseWriter`), which is used to write the HTTP response.
        - `r` (an `*http.Request`), which represents the client HTTP request.
    - `fmt.Fprintf(w, "Hello, World! This is a basic web server!")`: This line sends a string back to the client as part of the HTTP response.

4. **Main Function**:
    - `func main()`: The main function, which gets called when you run the program.
    - `http.HandleFunc("/", handler)`: This tells the http package to handle all requests to the web root ("/") with the `handler` function.
    - `fmt.Println("Server is running on http://localhost:8080")`: Prints a message to the console indicating that the server is running.
    - `http.ListenAndServe(":8080", nil)`: This starts an HTTP server with a given address (":8080") and handler (nil implies using the default handler, which you've set up as `handler` function for the root path).

To run this server:

1. **Save the Code**: Save the code into a file with a `.go` extension, for example, `server.go`.
2. **Run the Code**:
    - Open your terminal.
    - Navigate to the directory containing your `server.go` file.
    - Run the command `go run server.go`.
3. **Access the Server**: Open a web browser and go to `http://localhost:8080`. You should see "Hello, World! This is a basic web server!" displayed.
