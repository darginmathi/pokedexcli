package main

import "fmt"

func CommandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for name := range cfg.pokeDex {
		fmt.Printf(" - %v\n", name)
	}
	return nil
}
