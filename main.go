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

	// Routes
	apiRoutes.HandleFunc("/search-pokemon", api.SearchPokemon).Methods("GET")
	apiRoutes.HandleFunc("/get-pokemon",api.GetPokemon).Methods("GET")
	apiRoutes.HandleFunc("/get-ranks-for-iv",api.GetRanksForIV).Methods("GET")
	apiRoutes.HandleFunc("/get-ranks-for-iv-evolutions", api.GetRanksForIVEvolutions).Methods("GET")

	// Start the HTTP server
	http.Handle("/", r)
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
