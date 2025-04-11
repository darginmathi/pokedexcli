package main

import (
	"encoding/json"
	"fmt"

	"github.com/darginmathi/pokedexcli/internal/pokeapi"
)

func CommandMap(cfg *config) error {
	var locationResp pokeapi.LocationArea
	if cfg.nextLocationsURL == nil {
		// If no next URL, fetch the first page directly
		resp, err := cfg.pokeapiClient.ListLocations(nil)
		if err != nil {
			return err
		}
		locationResp = resp

		data, err := json.Marshal(locationResp)
		if err != nil {
			return err
		}

		Cache.Add("start", data)
		//check if in cache
	} else if data, ok := Cache.Get(*cfg.nextLocationsURL); !ok {

		resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
		if err != nil {
			return err
		}
		locationResp = resp

		data, err := json.Marshal(locationResp)
		if err != nil {
			return err
		}
		Cache.Add(*cfg.nextLocationsURL, data)

	} else {

		if err := json.Unmarshal(data, &locationResp); err != nil {
			return err
		}
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil

}

func CommandMapb(cfg *config) error {
	var locationResp pokeapi.LocationArea

	if cfg.prevLocationsURL == nil {
		// If no next URL, fetch the first page directly
		resp, err := cfg.pokeapiClient.ListLocations(nil)
		if err != nil {
			return err
		}
		locationResp = resp

	} else if data, t := Cache.Get(*cfg.prevLocationsURL); !t {

		resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
		if err != nil {
			return err
		}
		locationResp = resp

		data, err := json.Marshal(locationResp)
		if err != nil {
			return err
		}
		Cache.Add(*cfg.prevLocationsURL, data)
	} else {

		if err := json.Unmarshal(data, &locationResp); err != nil {
			return err
		}
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil

}
