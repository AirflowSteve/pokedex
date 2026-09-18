package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func commandCatch(conf *config, parameters []string) error {
	if len(parameters) < 1 {
		fmt.Println("what pokemon should I try to catch?")
		return nil
	}

	pokeTarget := strings.ToLower(parameters[0])
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

	return nil
}
