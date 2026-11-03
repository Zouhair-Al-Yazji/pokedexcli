package main

import (
	"fmt"
	"sort"
)

func commandHelp(cfg *config, args []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	names := make([]string, 0, len(cfg.commands))
	for name := range cfg.commands {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		cmd := cfg.commands[name]
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
