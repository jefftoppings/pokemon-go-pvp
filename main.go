package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/juju/ratelimit"

	"github.com/jefftoppings/pokemon-go-pvp/internal/api"
)

func main() {
	r := mux.NewRouter()
	apiRoutes := r.PathPrefix("/api").Subrouter()
	apiRoutes.Use(rateLimitMiddleware)
	apiRoutes.Use(corsMiddleware)

	// Routes
	apiRoutes.HandleFunc("/search-pokemon", api.SearchPokemon).Methods("GET")
	apiRoutes.HandleFunc("/get-pokemon",api.GetPokemon).Methods("GET")
	apiRoutes.HandleFunc("/get-ranks-for-iv",api.GetRanksForIV).Methods("GET")
	apiRoutes.HandleFunc("/get-ranks-for-iv-evolutions", api.GetRanksForIVEvolutions).Methods("GET")

	// Serve static files (for the frontend)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("dist")))

	http.Handle("/", r)

	// Start the HTTP server
	fmt.Println("Server is running on :8000")
	http.ListenAndServe(":8000", nil)
}

func rateLimitMiddleware(handler http.Handler) http.Handler {
	bucket := ratelimit.NewBucket(time.Second, 10) // 10 requests per second
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if bucket.TakeAvailable(1) > 0 {
			handler.ServeHTTP(w, r)
		} else {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		}
	})
}

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // List of allowed origins
        allowedOrigins := []string{
            "https://ionic-rank-checker.netlify.app/",
            "http://localhost:8100",
            "http://localhost:4200",
        }

        origin := r.Header.Get("Origin")
        for _, allowedOrigin := range allowedOrigins {
            if origin == allowedOrigin {
                // Allow requests from the specific origin
                w.Header().Set("Access-Control-Allow-Origin", origin)
                break
            }
        }

        // Allow other required headers
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            // Preflight request, respond with success
            w.WriteHeader(http.StatusOK)
            return
        }

        // Continue with the next handler
        next.ServeHTTP(w, r)
    })
}