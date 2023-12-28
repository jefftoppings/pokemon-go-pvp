package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jefftoppings/pokemon-go-pvp/internal/api"
)

type Config struct {
	APIKey string `json:"api_key"`
}

var config Config

func main() {
	err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/api/search-pokemon", AuthMiddleware(api.SearchPokemon)).Methods("GET")
	r.HandleFunc("/api/get-ranks-for-iv", AuthMiddleware(api.GetRanksForIV)).Methods("GET")
	r.HandleFunc("/api/get-ranks-for-iv-evolutions", AuthMiddleware(api.GetRanksForIVEvolutions)).Methods("GET")

	// Start the HTTP server
	http.Handle("/", r)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

// AuthMiddleware is a middleware function that checks for a valid API key in the request headers
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKeyFromRequest := r.Header.Get("X-API-Key")

		if apiKeyFromRequest != config.APIKey {
			http.Error(w, "Unauthorized. You need a valid api key to call this.", http.StatusUnauthorized)
			return
		}

		// Call the next handler if authentication is successful
		next.ServeHTTP(w, r)
	}
}

func loadConfig(filePath string) error {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read configuration file: %v", err)
	}

	if err := json.Unmarshal(fileContent, &config); err != nil {
		return fmt.Errorf("failed to unmarshal configuration data: %v", err)
	}

	return nil
}
