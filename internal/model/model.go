package model

type PokedexEntry struct {
	ID            string `json:"id"`
	FormID        string `json:"form_id"`
	PokedexNumber int64  `json:"pokedex_number"`
	Generation    int64  `json:"generation"`
	Name          string `json:"name"`
	Stats         struct {
		Attack  int64 `json:"attack"`
		Defense int64 `json:"defense"`
		Stamina int64 `json:"stamina"`
	} `json:"stats"`
	PrimaryType   string `json:"primary_type"`
	SecondaryType string `json:"secondary_type"`
	ImageUrl      string `json:"image_url"`
	Evolutions    []struct {
		ID string `json:"id"`
	} `json:"evolutions"`
}
