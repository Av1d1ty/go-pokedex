package main

import "errors"

func callbackInspect(cfg *config, args ...string) error {
    if len(args) == 0 {
        return errors.New("You must provide a pokemon name.")
    }

    pokemonName := args[0]
    pokemon, ok := cfg.caughtPokemon[pokemonName]
    if !ok {
        return errors.New("You don't have that pokemon.")
    }
    println("Name:", pokemon.Name)
    println("Height:", pokemon.Height)
    println("Weight:", pokemon.Weight)
    for _, stat := range pokemon.Stats {
        println(stat.Stat.Name, ":", stat.BaseStat)
    }
    for _, typ := range pokemon.Types {
        println("Type :", typ.Type.Name)
    }
    return nil
}
