package main

import (
	"errors"
	"fmt"
)

func explore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location given with command")
	}
	locationName := args[0]
	locationData, err := cfg.pokeapiClient.GetLocationData(locationName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", locationName)
	for _, pokemonEncounter := range locationData.PokemonEncounters {
		pokemonName := pokemonEncounter.Pokemon.Name
		fmt.Printf("- %s\n", pokemonName)
	}
	return nil
}
