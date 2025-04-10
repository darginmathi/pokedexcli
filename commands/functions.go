package commands

import (
	"fmt"
	"os"
)

// cli functions

func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range Commands {
		fmt.Printf("%v: %v\n", cmd.Name, cmd.Description)
	}
	return nil
}
