package commands

// init()

var Commands map[string]cliCommand

// cli command struct and commands

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func init() {
	Commands = map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
	}
}
