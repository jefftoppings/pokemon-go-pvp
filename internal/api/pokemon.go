package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jefftoppings/pokemon-go-pvp/internal/pokemon"
)

// SearchPokemon handles requests to /api/search-pokemon
func SearchPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pageSizeStr := r.URL.Query().Get("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		http.Error(w, "Invalid 'pageSize' parameter", http.StatusBadRequest)
		return
	}

	if pageSize == 0 {
		http.Error(w, "pageSize is required", http.StatusBadRequest)
		return
	}
	name := r.URL.Query().Get("name")

	results, err := pokemon.SearchPokemon(name, pageSize)
	if err != nil {
		errorMsg := fmt.Sprintf("Error searching pokemon: %v", err)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// GetPokemon handles requests to /api/get-pokemon
func GetPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	results, err := pokemon.GetPokemon(id)
	if err != nil {
		errorMsg := fmt.Sprintf("Error getting pokemon: %v", err)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}