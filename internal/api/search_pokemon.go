package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jefftoppings/pokemon-go-pvp/internal/model"
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

	name := r.URL.Query().Get("name")

	fmt.Printf("SearchPokemon name %s and page size %+v\n", name, pageSize)

	if pageSize == 0 {
		http.Error(w, "pageSize is required", http.StatusBadRequest)
		return
	}

	// TODO complete this
	results := []model.PokedexEntry{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
