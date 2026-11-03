package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Zouhair-Al-Yazji/pokedexcli/internal/pokeapi"
)

type config struct {
	commands      map[string]cliCommand
	nextURL       *string
	previousURL   *string
	caughtPokemon map[string]pokeapi.Pokemon
	pokeapiClient pokeapi.Client
}

type cliCommand struct {
	callback    func(*config, []string) error
	name        string
	description string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List the Pokemon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Show details of a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all the Pokemon you have caught",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]

		cmd, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(cfg, args); err != nil {
			fmt.Println(err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
}
