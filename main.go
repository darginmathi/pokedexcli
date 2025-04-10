package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/commands"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			cleaninput := cleanInput(input)
			if cmd, ok := commands.Commands[cleaninput[0]]; ok {
				if err := cmd.Callback(); err != nil {
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
