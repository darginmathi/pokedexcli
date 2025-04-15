package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonInfo(name string) (PokemonInfo, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		var pokemon PokemonInfo
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return PokemonInfo{}, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	{
		if err != nil {
			return PokemonInfo{}, err
		}
	}

	res, err := c.httpClient.Do(req)
	{
		if err != nil {
			return PokemonInfo{}, err
		}
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	{
		if err != nil {
			return PokemonInfo{}, err
		}
	}

	var pokemon PokemonInfo

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return PokemonInfo{}, err
	}

	c.cache.Add(url, data)

	return pokemon, nil
}
