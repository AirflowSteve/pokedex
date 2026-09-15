package main

import (
	"github.com/AirflowSteve/pokedex/internal/pokeapi"
)

func (conf *config) ListingPokemons(url string) (pokeapi.LocationInfo, error) {
	data, ok := conf.cache.Get(url)
	if ok {
		locationInfo, err := pokeapi.DecipherLocationInfo(data)
		if err != nil {
			return pokeapi.LocationInfo{}, err
		}
		return locationInfo, nil

	}

	data, err := pokeapi.Request(url, conf.PokeClient)
	if err != nil {
		return pokeapi.LocationInfo{}, err
	}

	conf.cache.Add(url, data)

	locationInfo, err := pokeapi.DecipherLocationInfo(data)
	if err != nil {
		return pokeapi.LocationInfo{}, err
	}

	return locationInfo, nil

}

// func (conf *config) ListingPokemonsCache(url string) (pokeapi.LocationInfo, error) {
// 	data, ok := conf.cache.Get(url)
// 	if ok {
// 		locationInfo, err := pokeapi.DecipherLocationInfo(data)
// 		if err != nil {
// 			return pokeapi.LocationInfo{}, err
// 		}
// 		return locationInfo, nil

// 	}

// 	return pokeapi.LocationInfo{}, nil

// }
