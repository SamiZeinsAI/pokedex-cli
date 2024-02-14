package main

import (
	"time"

	"github.com/SamiZeinsAI/pokedex-cli/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	repl(cfg)
}
