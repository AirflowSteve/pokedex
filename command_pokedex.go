package main

import "fmt"

func commandPokedex(conf *config, args []string) error {
	fmt.Println("Your Pokedex:")
	for pokemon := range conf.Pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}
