package main

import (
	"fmt"
)

func showLocations(cfg *config, pageURL *string) error {
	data, err := cfg.pokeapiClient.ListLocationAreas(pageURL)
	if err != nil {
		return err
	}

	cfg.nextURL = data.Next
	cfg.previousURL = data.Previous

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMap(cfg *config, args []string) error {
	return showLocations(cfg, cfg.nextURL)
}
