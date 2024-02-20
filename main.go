package main

import "github.com/Av1d1ty/go-pokedex.git/internal/pokeapi"

type config struct {
    pokeapiClient pokeapi.Client
    nextLocationAreaURL *string
    prevLocationAreaURL *string
}

func main() {
    cfg := config{
        pokeapiClient: pokeapi.NewClient(),
    }
    startRepl(&cfg)
}
