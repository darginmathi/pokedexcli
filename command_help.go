package main

import (
	"fmt"
)

func CommandHelp(cfg *config, arg ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	Commands := GetCommands()
	for _, cmd := range Commands {
		fmt.Printf("%v: %v\n", cmd.Name, cmd.Description)
	}
	return nil
}
