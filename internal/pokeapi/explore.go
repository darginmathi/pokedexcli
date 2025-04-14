package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(Location string) (ExploreLocation, error) {
	url := baseURL + "/location-area/" + Location

	if val, ok := c.cache.Get(url); ok {
		var exploreInfo ExploreLocation
		if err := json.Unmarshal(val, &exploreInfo); err != nil {
			return ExploreLocation{}, err
		}
		return exploreInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	{
		if err != nil {
			return ExploreLocation{}, err
		}
	}

	res, err := c.httpClient.Do(req)
	{
		if err != nil {
			return ExploreLocation{}, err
		}
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	{
		if err != nil {
			return ExploreLocation{}, err
		}
	}

	var locationInfo ExploreLocation

	if err := json.Unmarshal(data, &locationInfo); err != nil {
		return ExploreLocation{}, err
	}

	c.cache.Add(url, data)
	return locationInfo, nil
}
