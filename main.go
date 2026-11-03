package main

import (
	"time"

	"github.com/Zouhair-Al-Yazji/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		commands:      getCommands(),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
	}
	startRepl(cfg)
}
