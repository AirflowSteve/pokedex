package main

import (
	"fmt"
	"strings"
)

func commandExplore(conf *config, parameters []string) error {

	conf.Area = struct {
		Name              string
		URL               string
		PokemonsInTheArea []struct {
			Name string
			URL  string
		}
	}{}
	baseURL := defaultLocationsURL
	if len(parameters) < 1 {
		fmt.Println(">No areas listed")
		return nil
	}

	URLtoLookUp := baseURL + strings.ToLower(parameters[0])

	locationInfo, err := conf.ListingPokemons(URLtoLookUp)
	if err != nil {
		return err
	}

	conf.Area.Name = locationInfo.Location.Name
	conf.Area.URL = locationInfo.Location.URL

	fmt.Printf("Exploring %s...\n", locationInfo.Name)
	fmt.Println()
	fmt.Println("Found Pokemon:")

	for _, pokemonToMeet := range locationInfo.PokemonEncounters {
		conf.Area.PokemonsInTheArea = append(conf.Area.PokemonsInTheArea, struct {
			Name string
			URL  string
		}{
			Name: pokemonToMeet.Pokemon.Name,
			URL:  pokemonToMeet.Pokemon.URL,
		})
		fmt.Printf(" - %v\n", pokemonToMeet.Pokemon.Name)

	}
	fmt.Println()
	return nil
}

// func commandExploreCache(conf *config, parameters []string) error {
// 	baseURL := defaultLocationsURL
// 	if len(parameters) < 1 {
// 		fmt.Println(">No areas listed")
// 		return nil
// 	}

// 	for i := 0; i < len(parameters); i++ {
// 		URLtoLookUp := baseURL + parameters[i]

// 		locationInfo, err := conf.ListingPokemonsCache(URLtoLookUp)
// 		if err != nil {
// 			return err
// 		}

// 		fmt.Printf("Exploring %s...\n", locationInfo.Name)
// 		fmt.Println("Found Pokemon:")
// 		for _, pokemonToMeet := range locationInfo.PokemonEncounters {
// 			fmt.Printf(" - %v\n", pokemonToMeet.Pokemon.Name)
// 		}
// 		println()
// 	}

// 	return nil
// }
