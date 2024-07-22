package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *Config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("your Pokedex is empty")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Println(" -", pokemon.Name)
	}

	return nil
}
