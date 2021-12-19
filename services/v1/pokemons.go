package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PokemonResult struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type PokemonResponse struct {
	Count    int             `json:"count"`
	Next     string          `json:"next"`
	Previous *string         `json:"previous"`
	Results  []PokemonResult `json:"results"`
}

func Get() PokemonResponse {
	pokemons, err := http.Get("https://pokeapi.co/api/v2/pokemon/")

	if err != nil {
		fmt.Println(err)
	}

	p, err := ioutil.ReadAll(pokemons.Body)

	if err != nil {
		fmt.Println(err)
	}
	var pokemonResponse PokemonResponse
	json.Unmarshal(p, &pokemonResponse)
	return pokemonResponse
}
