package commands

import (
	"fmt"
	"os"

	"github.com/darginmathi/pokedexcli/internal/api/location"
	"github.com/darginmathi/pokedexcli/structs"
)

// cli functions

func CommandExit(cfg *structs.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(cfg *structs.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range Commands {
		fmt.Printf("%v: %v\n", cmd.Name, cmd.Description)
	}
	return nil
}

func CommandMap(cfg *structs.Config) error {
	if err := location.GetLocation(cfg, "next"); err != nil {
		return err
	}
	return nil
}

func CommandMapb(cfg *structs.Config) error {
	if err := location.GetLocation(cfg, "prev"); err != nil {
		return err
	}
	return nil
}
