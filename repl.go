package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		inputCommand, exists := getCommands()[commandName]
		if !exists {
			fmt.Printf("'%s' command not found\n", words[0])
			continue
		}

		err := inputCommand.callback(cfg, words[1:]...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower((text))
	words := strings.Fields(lowercaseText)
	return words
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    help_cmd,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    exit_cmd,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations from PokeAPI",
			callback:    map_cmd,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations from PokeAPI",
			callback:    mapb_cmd,
		},
		"explore": {
			name:        "explore",
			description: "Displays all pokemon in a given area",
			callback:    explore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch given pokemon",
			callback:    catch,
		},
		"inspect": {
			name:        "inspect",
			description: "Provides details of a given pokemon if it has been caught before",
			callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows pokemon in users pokedex",
			callback:    pokedex,
		},
	}
}
