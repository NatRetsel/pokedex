package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Get informaiton on the Pokemon
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
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
	if err!= nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err!= nil {
		return Pokemon{}, err
	}
	pokemonResp := Pokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err!= nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, data)
	return pokemonResp, nil
}