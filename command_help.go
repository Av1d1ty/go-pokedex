package main

func callbackHelp(cfg *config, args ...string) error {
    availableCommands := getCommands()
    for _, cmd := range availableCommands {
        println(cmd.name, "-", cmd.description)
    }
    return nil
}
