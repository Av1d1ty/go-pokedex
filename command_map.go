package main

import (
	"errors"
)

func callbackMap(cfg *config) error {
    resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreaURL)
    if err != nil {
        return err
    }
    println("Location Areas:")
    for _, area := range resp.Results {
        println(" - ", area.Name)
    }
    cfg.nextLocationAreaURL = resp.Next
    cfg.prevLocationAreaURL = resp.Previous
    return nil
}

func callbackMapBack(cfg *config) error {
    if cfg.prevLocationAreaURL == nil {
        return errors.New("You are on the first page.")
    }
    resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationAreaURL)
    if err != nil {
        return err
    }
    println("Location Areas:")
    for _, area := range resp.Results {
        println(" - ", area.Name)
    }
    cfg.nextLocationAreaURL = resp.Next
    cfg.prevLocationAreaURL = resp.Previous
    return nil
}
