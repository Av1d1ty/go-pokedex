package main

import (
	"bufio"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scaner := bufio.NewScanner(os.Stdin)
	for {
        print("> ")
		scaner.Scan()
		input := cleanInput(scaner.Text())
		if len(input) == 0 {
			continue
		}

		command := input[0]
        args := []string{}
        if len(input) > 1 {
            args = input[1:]
        }

        availableCommands := getCommands()
        cmd, ok := availableCommands[command]
        if !ok {
            println("Unknown command:", command)
            continue
        }
        err := cmd.callback(cfg, args...)
        if err != nil {
            println("Error:", err.Error())
        }
	}
}

type command struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func getCommands() map[string]command {
	return map[string]command{
        "map": {
            name:        "map",
            description: "shows location areas",
            callback:    callbackMap,
        },
        "map-back": {
            name:        "map-back",
            description: "shows previous location areas",
            callback:    callbackMapBack,
        },
        "explore": {
            name:        "explore {location-area}",
            description: "explore a location area",
            callback:    callbackExplore,
        },
        "catch": {
            name:        "catch {pokemon}",
            description: "catch a pokemon",
            callback:    callbackCatch,
        },
        "inspect": {
            name:        "inspect {pokemon}",
            description: "inspect a caught pokemon",
            callback:    callbackInspect,
        },
        "pokedex": {
            name:        "pokedex",
            description: "show caught pokemon",
            callback:    callbackPokedex,
        },
		"help": {
			name:        "help",
			description: "shows this help",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "exits the program",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	return strings.Fields(lowered)
}
