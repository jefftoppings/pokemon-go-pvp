package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"

    "github.com/jefftoppings/pokemon-go-pvp/internal/api"
)

func main() {
    r := mux.NewRouter()

    // Routes
	r.HandleFunc("/api/search-pokemon", api.SearchPokemon).Methods("GET")

    // Start the HTTP server
    http.Handle("/", r)
    fmt.Println("Server is running on :8080")
    http.ListenAndServe(":8080", nil)
}
