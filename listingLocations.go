package main

import (
	"github.com/AirflowSteve/pokedex/internal/pokeapi"
)

func (conf *config) listLocations(url string) (pokeapi.PokeAPIResponse, error) {
	pokeCache := &conf.cache

	val, ok := pokeCache.Get(url)
	if ok {
		response, err := pokeapi.Decipher(val)
		if err != nil {
			return pokeapi.PokeAPIResponse{}, err
		}

		return response, nil
	}
	data, err := pokeapi.Request(url, conf.PokeClient)
	if err != nil {
		return pokeapi.PokeAPIResponse{}, err
	}

	pokeCache.Add(url, data)

	response, err := pokeapi.Decipher(data)
	if err != nil {
		return pokeapi.PokeAPIResponse{}, err
	}

	return response, nil
}
