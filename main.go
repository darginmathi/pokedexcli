package main

import (
	"time"

	"github.com/darginmathi/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	// init cfg
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
