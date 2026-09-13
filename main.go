package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input := ""
		if scanner.Scan() {
			input = scanner.Text()
		} else {
			fmt.Errorf("uh-uh, didn't get it\n")
		}
		clean_input := cleanInput(input)
		if len(clean_input) != 0 {
			fmt.Printf("Your command was: %s\n", clean_input[0])
		} else {
			fmt.Println("Don't ignore me, please")
		}
	}
}
