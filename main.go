package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    // Define your API routes and handlers
    r.HandleFunc("/api/hello", HelloHandler).Methods("GET")

    // Start the HTTP server
    http.Handle("/", r)
    fmt.Println("Server is running on :8080")
    http.ListenAndServe(":8080", nil)
}

// HelloHandler handles requests to /api/hello
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, API!")
}
