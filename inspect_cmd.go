package main

import (
	"errors"
	"fmt"
)

func inspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("inspect requires a pokemon name as an argument")
	}
	pokemonName := args[0]
	if pokemon, exists := cfg.caughtPokemon[pokemonName]; exists {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, statData := range pokemon.Stats {
			fmt.Printf("	-%s: %v\n", statData.Stat.Name, statData.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeData := range pokemon.Types {
			fmt.Printf("	-%s\n", typeData.Type.Name)
		}
	} else {
		fmt.Printf("A %s has not been caught yet\n", pokemonName)
		return nil
	}
	return nil
}
