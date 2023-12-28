package get_ranks_for_iv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jefftoppings/pokemon-go-pvp/internal/model"
)

const (
	NOT_FOUND = "not found"
)

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
