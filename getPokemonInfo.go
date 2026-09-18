package main

import (
	"github.com/AirflowSteve/pokedex/internal/pokeapi"
)

func getPokemonInfo(conf *config, url string) (pokeapi.PokemonDesc, error) {
	pokeCache := &conf.cache

	val, ok := pokeCache.Get(url)
	if ok {

		response, err := pokeapi.DecipherPokemonInfo(val)
		if err != nil {
			return pokeapi.PokemonDesc{}, err
		}

		return response, nil
	}

	data, err := pokeapi.Request(url, conf.PokeClient)
	if err != nil {
		return pokeapi.PokemonDesc{}, err
	}

	pokeCache.Add(url, data)

	response, err := pokeapi.DecipherPokemonInfo(data)
	if err != nil {
		return pokeapi.PokemonDesc{}, err
	}

	return response, nil

}
