package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/jefftoppings/pokemon-go-pvp/internal/model"
)

const (
	pokedexPath = "internal/assets/pokedex.json"
)

var (
	allPokemon     []*model.Pokemon
)

func init() {
	loadPokemonData()
}

func loadPokemonData() {
	pokedexContent, err := ioutil.ReadFile(pokedexPath)
	if err != nil {
		fmt.Printf("failed to read pokedex JSON pokedex data: %v", err)
		return
	}

	if err := json.Unmarshal(pokedexContent, &allPokemon); err != nil {
		fmt.Printf("failed to unmarshal JSON pokedex content: %v", err)
		return
	}
}

func SearchPokemon(name string, pageSize int) ([]*model.Pokemon, error) {
	results := []*model.Pokemon{}
	nameLower := strings.ToLower(name)

	// Concurrently search for pokemon by name
	resultChan := make(chan *model.Pokemon, len(allPokemon))
	var wg sync.WaitGroup

	for _, pokemon := range allPokemon {
		wg.Add(1)
		go func(pokemon *model.Pokemon) {
			defer wg.Done()

			pokemonNameLower := strings.ToLower(pokemon.Names.English)
			if strings.Contains(pokemonNameLower, nameLower) {
				resultChan <- pokemon
			}
		}(pokemon)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results
	for result := range resultChan {
		if len(results) < pageSize {
			results = append(results, result)
		} else {
			break
		}
	}
	return results, nil
}

func GetPokemon(id string) (*model.Pokemon, error) {
	// find the specific pokemon
	// split allPokemon into chunks of 100 to parallelize the process
	const chunkSize = 100
	numChunks := (len(allPokemon) + chunkSize - 1) / chunkSize
	resultChan := make(chan *model.Pokemon, numChunks)
	var wg sync.WaitGroup

	for i := 0; i < numChunks; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > len(allPokemon) {
			end = len(allPokemon)
		}

		wg.Add(1)
		go findPokemonInChunk(allPokemon[start:end], id, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		if result != nil {
			return result, nil
		}
	}

	return nil, fmt.Errorf("pokemon with ID %s not found", id)
}

func findPokemonInChunk(chunk []*model.Pokemon, id string, wg *sync.WaitGroup, resultChan chan<- *model.Pokemon) {
	defer wg.Done()

	for _, pokemon := range chunk {
		if strings.EqualFold(pokemon.ID, id) {
			resultChan <- pokemon
			return
		}
	}

	resultChan <- nil
}
