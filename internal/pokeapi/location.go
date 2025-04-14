package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// checking for cache availablity
	if val, ok := c.cache.Get(url); ok {
		var locationResp LocationArea
		if err := json.Unmarshal(val, &locationResp); err != nil {
			return LocationArea{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	var locationResp LocationArea

	if err := json.Unmarshal(data, &locationResp); err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, data)
	return locationResp, nil
}
