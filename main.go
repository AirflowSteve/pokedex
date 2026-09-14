package main

func main() {
	conf := config{
		CommandRegistry: map[string]cliCommands{
			"exit": {
				name:        "exit",
				description: "Exit the pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
			"map": {
				name: "map",
				description: "Shows a map of the pokemon world",
				callback: commandMap,
			},
		},
	}
	startRepl(&conf)
}
