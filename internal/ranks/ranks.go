package ranks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jefftoppings/pokemon-go-pvp/internal/model"
	"github.com/jefftoppings/pokemon-go-pvp/internal/pokemon"
)

const (
	NOT_FOUND = "not found"
)

// pokemonEvolutions represents the structure of the evolutions.
type pokemonEvolutions struct {
	ID         string
	Evolutions []*pokemonEvolutions
}

func GetRanksForIV(id string, attack int, defense int, stamina int) (*model.GetRanksForIVResponse, error) {
	// id should be lower case for the file names
	id = strings.ToLower(id)

	// determine files to lookup data from
	greatLeagueFile := fmt.Sprintf("internal/assets/data/great/%s.json", id)
	ultraLeagueFile := fmt.Sprintf("internal/assets/data/ultra/%s.json", id)

	var greatLeagueDataMap map[string]model.PokemonIVData
	var ultraLeagueDataMap map[string]model.PokemonIVData

	greatLeagueData, err := ioutil.ReadFile(greatLeagueFile)
	if err != nil {
		return nil, fmt.Errorf("ID %s %s in great league files: %v", id, NOT_FOUND, err)
	}
	err = json.Unmarshal(greatLeagueData, &greatLeagueDataMap)
	if err != nil {
		return nil, fmt.Errorf("ID %s %s not found in ultra league files: %v", id, NOT_FOUND, err)
	}

	ultraLeagueData, err := ioutil.ReadFile(ultraLeagueFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(ultraLeagueData, &ultraLeagueDataMap)
	if err != nil {
		return nil, err
	}

	// get the rank for the given IVs
	ivs := fmt.Sprintf("%v/%v/%v", attack, defense, stamina)

	greatValue, exists := greatLeagueDataMap[ivs]
	if !exists {
		return nil, fmt.Errorf("IVs %s %s in great league data", ivs, NOT_FOUND)
	}
	ultraValue, exists := ultraLeagueDataMap[ivs]
	if !exists {
		return nil, fmt.Errorf("IVs %s %s not found in ultra league data", ivs, NOT_FOUND)
	}

	return &model.GetRanksForIVResponse{
		GreatLeagueRank: greatValue,
		UltraLeagueRank: ultraValue,
	}, nil
}

func GetRanksForIVEvolutions(id string, attack int, defense int, stamina int) (*model.GetRanksForIVEvolutionsResponse, error) {
	id = strings.ToUpper(id)
	// determine the evolutions to lookup
	evolutions, err := getPokemonEvolutions(id)
	if err != nil {
		return nil, err
	}
	evolutionIDs := getEvolutionIDs(evolutions)

	// build up maps for response
	rankForEvolutions := map[string]model.GetRanksForIVResponse{}
	for _, evolutionID := range evolutionIDs {
		ranks, err := GetRanksForIV(evolutionID, attack, defense, stamina)
		if err != nil {
			return nil, err
		}
		rankForEvolutions[evolutionID] = *ranks
	}

	return &model.GetRanksForIVEvolutionsResponse{
		Evolutions:        evolutionIDs,
		RankForEvolutions: rankForEvolutions,
	}, nil
}

// getEvolutionIDs returns a slice of IDs of the evolutions in the given pokemonEvolutions.
func getEvolutionIDs(evolution *pokemonEvolutions) []string {
	var evolutionIDs []string
	// Add the ID of the current evolution
	evolutionIDs = append(evolutionIDs, evolution.ID)

	// Add IDs of sub-evolutions
	for _, subEvolution := range evolution.Evolutions {
		evolutionIDs = append(evolutionIDs, getEvolutionIDs(subEvolution)...)
	}
	return evolutionIDs
}

// GetPokemonEvolutions retrieves the evolution chain for a given Pokemon ID.
func getPokemonEvolutions(id string) (*pokemonEvolutions, error) {
	pokemon, err := pokemon.GetPokemon(id)
	if err != nil {
		return nil, fmt.Errorf("ID %s %s: %v", id, NOT_FOUND, err)
	}

	evolutions := make([]*pokemonEvolutions, len(pokemon.Evolutions))
	for i, evolution := range pokemon.Evolutions {
		evolution, err := getPokemonEvolutions(evolution.ID)
		if err != nil {
			return nil, err
		}
		evolutions[i] = evolution
	}

	return &pokemonEvolutions{
		ID:         id,
		Evolutions: evolutions,
	}, nil
}
