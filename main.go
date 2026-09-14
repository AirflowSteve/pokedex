package main

import "net/http"

func main() {
	conf := &config{
		CommandRegistry: getCommands(),
		Previous:        "",
		Next:            "https://pokeapi.co/api/v2/location-area/",
		PokeClient:      http.Client{},
	}
	startRepl(conf)
}
