package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) ListLocations(pageURL *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, nil
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, nil
	}

	var areas LocationArea

	if err := json.Unmarshal(data, &areas); err != nil {
		return LocationArea{}, err
	}

	return areas, nil
}
