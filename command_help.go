package main

func callbackHelp(cfg *config) error {
    availableCommands := getCommands()
    for _, cmd := range availableCommands {
        println(cmd.name, "-", cmd.description)
    }
    return nil
}
