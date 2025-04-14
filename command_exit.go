package main

import (
	"fmt"
	"os"
)

func CommandExit(cfg *config, arg ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
