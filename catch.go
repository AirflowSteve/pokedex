package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func commandCatch(conf *config, parameters []string) error {
	if conf.Area.Name == "" {

		fmt.Println("you're not in an area yet")

		return nil
	}
	if len(parameters) < 1 {

		fmt.Println("what pokemon should I try to catch?")

		return nil
	}

	pokeTarget := strings.ToLower(parameters[0])

	ok := checkIfInArea(conf, pokeTarget)
	if !ok {

		fmt.Println("No such pokemon in the area")

		return nil
	}

	PokemonURL := defaultPokeURL + "/pokemon/" + pokeTarget

	fmt.Printf("Throwing a Pokeball at %s...\n", pokeTarget)

	PokemonInfo, err := getPokemonInfo(conf, PokemonURL)
	if err != nil {
		return err
	}

	baseLevel := PokemonInfo.BaseExperience
	if baseLevel == 0 {
		// fmt.Println("I don't know such a Pokemon")
		return fmt.Errorf("Unknown Pokemon '%v'", pokeTarget)

	}
	res := rand.Intn(baseLevel)

	if res > 40 {
		fmt.Printf("%s escaped!\n", pokeTarget)

		return nil
	}
	fmt.Printf("%s was caught!\n", pokeTarget)
	conf.Pokedex[PokemonInfo.Name] = PokemonInfo
	err = saveFunction(conf)
	if err != nil {
		return err
	}

	return nil
}

func checkIfInArea(conf *config, pokeName string) bool {
	for _, pokemon := range conf.Area.PokemonsInTheArea {
		if strings.ToLower(pokemon.Name) == strings.ToLower(pokeName) {
			return true
		}
	}
	return false
}
