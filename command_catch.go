package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("need to enter a Pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.PokemonInfo(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	if rand.Intn(10) > 400/pokemon.BaseExperience {
		fmt.Printf("%v escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%v was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	if _, ok := cfg.pokeDex[pokemon.Name]; !ok {
		cfg.pokeDex[pokemon.Name] = pokemon
	}

	return nil
}
