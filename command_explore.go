package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("need to enter location name to explore")
	}

	name := args[0]
	exploreResp, err := cfg.pokeapiClient.ExploreLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", exploreResp.Name)
	fmt.Print("Found Pokemon:\n")
	for _, enc := range exploreResp.PokemonEncounters {
		fmt.Printf(" - %v\n", enc.Pokemon.Name)
	}

	return nil
}
