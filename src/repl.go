package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/janmmiranda/repl-dex/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsUrl     *string
	previousLocationsUrl *string
	pokemonBox           pokeBox
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		argument := ""
		if len(words) > 1 {
			argument = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, argument)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unkown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get next locations, 20 at a time.",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous locations, 20 at a time.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Find pokemon given an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon, given its name!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon given it's name",
			callback:    commandInspect,
		},
	}
}
