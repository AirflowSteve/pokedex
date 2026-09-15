package main

import (
	"fmt"
)

func commandExplore(conf *config, parameters []string) error {
	baseURL := defaultLocationsURL
	if len(parameters) < 1 {
		fmt.Println(">No areas listed")
		return nil
	}

	for i := 0; i < len(parameters); i++ {
		URLtoLookUp := baseURL + parameters[i]

		locationInfo, err := conf.ListingPokemons(URLtoLookUp)
		if err != nil {
			return err
		}

		fmt.Printf("Exploring %s...\n", locationInfo.Name)
		fmt.Println("Found Pokemon:")
		for _, pokemonToMeet := range locationInfo.PokemonEncounters {
			fmt.Printf(" - %v\n", pokemonToMeet.Pokemon.Name)
		}
		println()
	}

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
