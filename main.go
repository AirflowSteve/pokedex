package main

func main() {
	conf := config{
		CommandRegistry: getCommands(),
		Previous:        "",
		Next:            "https://pokeapi.co/api/v2/location-area/",
	}
	startRepl(&conf)
}
