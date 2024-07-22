package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *Config, args ...string) error {
	// if cfg.NextLocationAreaURL == nil {
	// 	return errors.New("already on the last page")
	// }
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.NextLocationAreaURL)
	if err != nil {
		return err
	}

	cfg.NextLocationAreaURL = resp.Next
	cfg.PreviousLocationAreaURL = resp.Previous

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}

func commandMapb(cfg *Config, args ...string) error {
	if cfg.PreviousLocationAreaURL == nil {
		return errors.New("already on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.PreviousLocationAreaURL)
	if err != nil {
		return err
	}

	cfg.NextLocationAreaURL = resp.Next
	cfg.PreviousLocationAreaURL = resp.Previous

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}
