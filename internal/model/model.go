package model

type Pokemon struct {
	ID            string           `json:"id"`
	FormID        string           `json:"formId"`
	DexNr         int              `json:"dexNr"`
	Generation    int              `json:"generation"`
	Names         NameTranslations `json:"names"`
	Stats         PokemonStats     `json:"stats"`
	PrimaryType   PokemonType      `json:"primaryType"`
	SecondaryType PokemonType      `json:"secondaryType"`
	Assets        PokemonAssets    `json:"assets"`
	Evolutions    []Evolution      `json:"evolutions"`
}

type NameTranslations struct {
	English string `json:"English"`
}

type PokemonStats struct {
	Stamina int `json:"stamina"`
	Attack  int `json:"attack"`
	Defense int `json:"defense"`
}

type PokemonType struct {
	Type  string           `json:"type"`
	Names NameTranslations `json:"names"`
}

type PokemonAssets struct {
	Image      string `json:"image"`
	ShinyImage string `json:"shinyImage"`
}

type Evolution struct {
	ID     string `json:"id"`
	FormID string `json:"formId"`
}

type GetRanksForIVResponse struct {
	GreatLeagueRank PokemonIVData `json:"greatLeagueRank"`
	UltraLeagueRank PokemonIVData `json:"ultraLeagueRank"`
}

type GetRanksForIVEvolutionsResponse struct {
	Evolutions        []string                         `json:"evolutions"`
	RankForEvolutions map[string]GetRanksForIVResponse `json:"rankForEvolutions"`
}

type PokemonIVData struct {
	Ranks       map[string]int `json:"ranks"`
	Ivs         string         `json:"ivs"`
	Level       float64        `json:"level"`
	CP          int            `json:"cp"`
	StatProduct float64        `json:"statProduct"`
	Percent     float64        `json:"percent"`
	Stats       struct {
		Atk float64 `json:"atk"`
		Def float64 `json:"def"`
		Hp  float64 `json:"hp"`
	} `json:"stats"`
}
