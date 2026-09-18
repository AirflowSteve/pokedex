package main

import (
	"net/http"
	"time"

	"github.com/AirflowSteve/pokedex/internal/pokeapi"
	"github.com/AirflowSteve/pokedex/internal/pokecache"
)

func main() {
	conf := &config{
		CommandRegistry: getCommands(),
		Previous:        "",
		Next:            "https://pokeapi.co/api/v2/location-area/",
		PokeClient:      http.Client{},
		cache:           pokecache.NewCache(5 * time.Second),
		Pokedex:         make(map[string]pokeapi.PokemonDesc),
	}
	startRepl(conf)
}
