package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/jefftoppings/pokemon-go-pvp/internal/ranks"
)

// GetRanksForIV handles requests to /api/get-ranks-for-iv
func GetRanksForIV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, attack, defense, stamina, err := validateAndGetRanksParams(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	results, err := ranks.GetRanksForIV(id, attack, defense, stamina)
	if err != nil {
		if strings.Contains(err.Error(), ranks.NOT_FOUND) {
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

// GetRanksForIVEvolutions handles requests to /api/get-ranks-for-iv-evolutions
func GetRanksForIVEvolutions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, attack, defense, stamina, err := validateAndGetRanksParams(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	results, err := ranks.GetRanksForIVEvolutions(id, attack, defense, stamina)
	if err != nil {
		if strings.Contains(err.Error(), ranks.NOT_FOUND) {
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

func validateAndGetRanksParams(params url.Values) (id string, attack int, defense int, stamina int, err error) {
	attackStr := params.Get("attack")
	attack, err = strconv.Atoi(attackStr)
	if err != nil || attack < 0 || attack > 15 {
		err = fmt.Errorf("Invalid 'attack' parameter. It must be between 0 and 15.")
		return
	}

	defenseStr := params.Get("defense")
	defense, err = strconv.Atoi(defenseStr)
	if err != nil || defense < 0 || defense > 15 {
		err = fmt.Errorf("Invalid 'defense' parameter. It must be between 0 and 15.")
		return
	}

	staminaStr := params.Get("stamina")
	stamina, err = strconv.Atoi(staminaStr)
	if err != nil || stamina < 0 || stamina > 15 {
		err = fmt.Errorf("Invalid 'stamina' parameter. It must be between 0 and 15.")
		return
	}

	id = params.Get("id")
	if id == "" {
		err = fmt.Errorf("Missing 'id' parameter")
		return
	}
	return
}
