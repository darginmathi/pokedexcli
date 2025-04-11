package main

import (
	"fmt"
	"time"

	"github.com/darginmathi/pokedexcli/internal/pokeapi"
	"github.com/darginmathi/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	Cache = pokecache.NewCache(5 * time.Minute)
	fmt.Println("cache initialised:", Cache)

	// init cfg
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}

var Cache *pokecache.Cache
