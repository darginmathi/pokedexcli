package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("need to enter a Pokemon name")
	}

	name := args[0]

	pokemon, ok := cfg.pokeDex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Printf("Stats: \n")
	stats := pokemon.Stats
	for _, stat := range stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types: \n")
	types := pokemon.Types
	for _, typ := range types {
		fmt.Printf("  - %v\n", typ.Type.Name)
	}

	return nil
}
