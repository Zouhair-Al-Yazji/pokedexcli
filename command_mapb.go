package main

import (
	"fmt"
)

func commandMapb(cfg *config, args []string) error {
	if cfg.previousURL == nil {
		fmt.Println("You're in the first page")
		return nil
	}
	return showLocations(cfg, cfg.previousURL)
}
