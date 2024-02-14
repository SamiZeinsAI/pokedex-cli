package main

import "fmt"

func help_cmd(cfg *config, args ...string) error {
	fmt.Printf("This is a Pokedex\nHere are all available commands:\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
