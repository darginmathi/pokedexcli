package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/darginmathi/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokeDex          map[string]pokeapi.PokemonInfo
}

func startRepl(cfg *config) {
	// waiting for user input
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()
		// loop starts after user input
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := GetCommands()[commandName]
		if exists {
			if err := command.Callback(cfg, args...); err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("unknown command")
			continue
		}

	}
}

// helper fn

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*config, ...string) error
}

// cli commands init()

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of next 20 location areas in the Pokemon world",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the names of previous 20 location areas in the Pokemon world",
			Callback:    CommandMapb,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"explore": {
			Name:        "explore",
			Description: "Lists all the pokemon located in the area",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Tries to catch a Pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a pokemon",
			Callback:    commandInspect,
		},
	}
}
