package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jefftoppings/pokemon-go-pvp/internal/get_ranks_for_iv"
)

// GetRanksForIV handles requests to /api/get-ranks-for-iv
func GetRanksForIV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	attackStr := r.URL.Query().Get("attack")
	attack, err := strconv.Atoi(attackStr)
	if err != nil || attack < 0 || attack > 15 {
		http.Error(w, "Invalid 'attack' parameter. It must be between 0 and 15.", http.StatusBadRequest)
		return
	}

	defenseStr := r.URL.Query().Get("defense")
	defense, err := strconv.Atoi(defenseStr)
	if err != nil || defense < 0 || defense > 15 {
		http.Error(w, "Invalid 'defense' parameter. It must be between 0 and 15.", http.StatusBadRequest)
		return
	}

	staminaStr := r.URL.Query().Get("stamina")
	stamina, err := strconv.Atoi(staminaStr)
	if err != nil || stamina < 0 || stamina > 15 {
		http.Error(w, "Invalid 'defense' parameter. It must be between 0 and 15.", http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	fmt.Printf("ID: %+v, Attack: %v, Defense: %v, Stamina: %v\n", id, attack, defense, stamina)

	results, err := get_ranks_for_iv.GetRanksForIV(id, attack, defense, stamina)
	if err != nil {
		if strings.Contains(err.Error(), get_ranks_for_iv.NOT_FOUND) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		errorMsg := fmt.Sprintf("Error getting ranks: %v", err)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
