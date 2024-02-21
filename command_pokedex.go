package main

func callbackPokedex(cfg *config, args ...string) error {
    for _, pokemon := range cfg.caughtPokemon {
        callbackInspect(cfg, pokemon.Name)
        println("----------")
    }
    return nil
}
