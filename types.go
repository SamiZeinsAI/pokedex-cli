package main

import (
	"github.com/SamiZeinsAI/pokedex-cli/internal/pokeapi"
)

type config struct {
	caughtPokemon map[string]pokeapi.Pokemon
	pokeapiClient pokeapi.Client
	previous      *string
	next          *string
}

type command struct {
	name        string
	description string
	callback    func(*config, ...string) error
}
