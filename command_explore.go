package main

import (
	"errors"
)

func callbackExplore(cfg *config, args ...string) error {
    if len(args) == 0 {
        return errors.New("You must provide a location area name.")
    }
    locationAreaName := args[0]
    locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
    if err != nil {
        return err
    }

    println("Pokemons in", locationAreaName)
    for _, pokemon := range locationArea.PokemonEncounters {
        println(" - ", pokemon.Pokemon.Name)
    }

    return nil
}
