package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseUrl + "pokemon/" + pokemonName
	pokemonResp := Pokemon{}
	if val, exists := c.cache.Get(url); exists {
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	err = json.Unmarshal(body, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, body)
	return pokemonResp, nil
}
