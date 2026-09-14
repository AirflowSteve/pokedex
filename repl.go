package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommands struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	CommandRegistry map[string]cliCommands
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
			err := command.callback(conf)
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

func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for cmdName, value := range conf.CommandRegistry {
		fmt.Printf("%s: %s\n", cmdName, value.description)
	}
	return nil
}

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
