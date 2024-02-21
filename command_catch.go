package main

import (
    "errors"
    "math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
    if len(args) == 0 {
        return errors.New("You must provide a pokemon name.")
    }
    pokemonName := args[0]
    pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
    if err != nil {
        return err
    }

    threshold := 50
    randNum := rand.Intn(pokemon.BaseExperience)
    if randNum > threshold {
        println("You failed to catch", pokemon.Name)
        return nil
    }
    cfg.caughtPokemon[pokemon.Name] = pokemon
    println("You caught", pokemon.Name)
    return nil
}
