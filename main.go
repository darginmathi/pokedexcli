package main

import "github.com/darginmathi/pokedexcli/structs"

func main() {
	// init cfg
	cfg := &structs.Config{
		Next: "https://pokeapi.co/api/v2/location-area/",
	}
	startRepl(cfg)
}
