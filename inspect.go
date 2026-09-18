package main

import "fmt"

func commandInspect(conf *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("What Pokemon do you want to know about?")
	}
	pokeTarget := args[0]
	pokemon, ok := conf.Pokedex[pokeTarget]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Println()
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, PokeStat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", PokeStat.Stat.Name, PokeStat.BaseStat)
	}
	fmt.Println("Types:")
	for _, PokeType := range pokemon.Types {
		fmt.Printf("  - %s\n", PokeType.Type.Name)
	}
	fmt.Println()
	return nil
}
