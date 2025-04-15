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
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types: \n")
	for _, typ := range pokemon.Types {
		fmt.Printf("  - %v\n", typ.Type.Name)
	}

	return nil
}
