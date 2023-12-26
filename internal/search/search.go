package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jefftoppings/pokemon-go-pvp/internal/model"
)

func SearchPokemon(name string, pageSize int) ([]model.Pokemon, error) {
	results := []model.Pokemon{}

	pokedexPath := "internal/assets/pokedex.json"
	pokedexContent, err := ioutil.ReadFile(pokedexPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read pokedex JSON data: %v", err)
	}

	// Unmarshal the JSON data into a slice of Pokemon
	var allPokemon []model.Pokemon
	if err := json.Unmarshal(pokedexContent, &allPokemon); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %v", err)
	}

	// filter pokemon by name
	for _, pokemon := range allPokemon {
		nameLower := strings.ToLower(name)
		pokemonNameLower := strings.ToLower(pokemon.Names.English)
		if strings.Contains(pokemonNameLower, nameLower) {
			if len(results) < pageSize {
				results = append(results, pokemon)
			} else {
				break
			}
		}
	}

	return results, nil
}
