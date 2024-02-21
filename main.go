package main

import (
	"time"

	"github.com/Av1d1ty/go-pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
        caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
