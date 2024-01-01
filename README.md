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

Now, you can visit `http://localhost:8080` in your browser to see the server running.

## Understanding the Code

- **Package Declaration**: `package main` - Indicates that this file belongs to the main package.
- **Imports**:
    - `fmt`: Utilized for formatting and printing text.
    - `net/http`: Provides HTTP client and server implementations.
- **Handler Function**:
    - `func handler(w http.ResponseWriter, r *http.Request)`: Called for every HTTP request to the server.
    - `fmt.Fprintf(w, "Hello, World! This is a basic web server!")`: Writes the greeting message to the HTTP response.
- **Main Function**:
    - `func main()`: Entry point of the program.
    - `http.HandleFunc("/", handler)`: Directs the root URL to the handler function.
    - `fmt.Println("Server is running on http://localhost:8080")`: Indicates server is running.
    - `http.ListenAndServe(":8080", nil)`: Starts the HTTP server.

## Contributing

Contributions, issues, and feature requests are welcome! Feel free to check [issues page](link).

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.

## Authors

- **[Your Name](Your Github Profile Link)** - Initial work

## Acknowledgments

- Hat tip to anyone whose code was used
- Inspiration
- etc.

Replace placeholders and add additional sections as needed for your specific project. This README is structured to provide an overview, instructions on getting started, understanding the code, and other relevant information for anyone interested in your project.