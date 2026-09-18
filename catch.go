package main

import (
	"fmt"
)

func commandCatch(conf *config, parameters []string) error {
	if len(parameters) < 1 {
		fmt.Println("what pokemon should I try to catch?")
		return nil
	}

	pokeTarget := parameters[0]
	fmt.Printf("Throwing a pokeball at %s...\n", pokeTarget)
	return nil
}
