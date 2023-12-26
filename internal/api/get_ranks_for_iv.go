package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jefftoppings/pokemon-go-pvp/internal/get_ranks_for_iv"
)

// GetRanksForIV handles requests to /api/get-ranks-for-iv
func GetRanksForIV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	attackStr := r.URL.Query().Get("attack")
	attack, err := strconv.Atoi(attackStr)
	if err != nil {
		http.Error(w, "Invalid 'attack' parameter", http.StatusBadRequest)
		return
	}

	defenseStr := r.URL.Query().Get("defense")
	defense, err := strconv.Atoi(defenseStr)
	if err != nil {
		http.Error(w, "Invalid 'defense' parameter", http.StatusBadRequest)
		return
	}

	staminaStr := r.URL.Query().Get("stamina")
	stamina, err := strconv.Atoi(staminaStr)
	if err != nil {
		http.Error(w, "Invalid 'defense' parameter", http.StatusBadRequest)
		return
	}

	cpLimitStr := r.URL.Query().Get("cpLimit")
	cpLimit, err := strconv.Atoi(cpLimitStr)
	if err != nil {
		http.Error(w, "Invalid 'defense' parameter", http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")

	fmt.Printf("ID: %+v, Attack: %v, Defense: %v, Stamina: %v, CPLimit: %v\n", id, attack, defense, stamina, cpLimit)

	results, err := get_ranks_for_iv.GetRanksForIV(id, attack, defense, stamina, cpLimit)
	if err != nil {
		errorMsg := fmt.Sprintf("Error getting ranks: %v", err)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
