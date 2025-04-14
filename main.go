package main

import (
	"time"

	"github.com/darginmathi/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	// init cfg
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
