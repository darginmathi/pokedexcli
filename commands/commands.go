package commands

import "github.com/darginmathi/pokedexcli/structs"

// init()

var Commands map[string]structs.CliCommand

// cli commands init()

func init() {
	Commands = map[string]structs.CliCommand{
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
	}
}
