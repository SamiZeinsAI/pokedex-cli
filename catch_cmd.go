package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("a valid pokemon name must be provided with this command")
	}
	pokemonName := args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	randNumber := rand.Intn(200)
	if pokemonResp.BaseExperience > randNumber {
		fmt.Printf("You caught a %s!!!\n", pokemonName)
		fmt.Println("It has been added to your pokedex")
		cfg.caughtPokemon[pokemonName] = pokemonResp
	} else {
		fmt.Printf("%s broke free from your pokeball\n", pokemonName)
	}
	return nil
}
