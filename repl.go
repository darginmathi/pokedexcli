package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/darginmathi/pokedexcli/commands"
	"github.com/darginmathi/pokedexcli/structs"
)

func startRepl(cfg *structs.Config) {
	// waiting for user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		// loop starts after user input
		if scanner.Scan() {
			input := scanner.Text()
			cleaninput := cleanInput(input)
			if cmd, ok := commands.Commands[cleaninput[0]]; ok {
				if err := cmd.Callback(cfg); err != nil {
					fmt.Println("Error:", err)
				}
			} else {
				fmt.Println("Unknows command")
			}
		}
	}
}

// helper fn

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
