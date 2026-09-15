package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/AirflowSteve/pokedex/internal/pokeapi"
	"github.com/AirflowSteve/pokedex/internal/pokecache"
)

const defaultLocationsURL string = "https://pokeapi.co/api/v2/location-area/"

type cliCommands struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	CommandRegistry map[string]cliCommands
	Previous        string
	Next            string
	PokeClient      http.Client
	cache           pokecache.Cache
	Area            string
}

func startRepl(conf *config) {

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		command, ok := conf.CommandRegistry[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback(conf, words[1:])
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func commandHelp(conf *config, parameters []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for cmdName, value := range conf.CommandRegistry {
		fmt.Printf("%s: %s\n", cmdName, value.description)
	}
	return nil
}

func commandExit(conf *config, parameters []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(conf *config, parameters []string) error {
	pokeAPIMapURL := conf.Next

	response, err := conf.listLocations(pokeAPIMapURL)
	if err != nil {
		return err
	}

	conf.Previous = response.PreviousURL
	conf.Next = response.NextURL

	locations := response.Results

	for _, loc := range locations {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(conf *config, parameters []string) error {
	pokeAPIMapURL := conf.Previous
	if pokeAPIMapURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	response, err := conf.listLocations(pokeAPIMapURL)
	if err != nil {
		return err
	}

	conf.Previous = response.PreviousURL
	conf.Next = response.NextURL

	locations := response.Results

	for _, loc := range locations {
		fmt.Println(loc.Name)
	}
	return nil
}

func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows a map of the pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Goes back to the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows all the pokemons in the area",
			callback:    commandExplore,
		},
		"cache": {
			name:        "cache",
			description: "Returns cache",
			callback:    getCache,
		},
	}
}

func getCache(conf *config, parameters []string) error {

	for _, value := range conf.cache.Cache {
		response, err := pokeapi.Decipher(value.Val)
		if err != nil {
			return err
		}
		fmt.Println(response)
	}

	return nil
}
