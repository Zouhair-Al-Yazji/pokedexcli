package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: catch <pokemon-name>")
	}

	name := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	const catchThreshold = 50
	if rand.Intn(pokemon.BaseExperience+1) > catchThreshold {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.caughtPokemon[name] = pokemon
	return nil
}
