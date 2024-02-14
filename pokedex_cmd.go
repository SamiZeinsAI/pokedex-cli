package main

import "fmt"

func pokedex(cfg *config, args ...string) error {
	fmt.Println("Your pokedex:")
	for k := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", k)
	}
	return nil
}
