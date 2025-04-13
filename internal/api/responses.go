package api

type LocationAreaResponse struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type LocationAreaDetailsResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
