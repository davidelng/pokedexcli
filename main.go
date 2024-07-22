package main

import (
	"time"

	"github.com/davidelng/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient           pokeapi.Client
	NextLocationAreaURL     *string
	PreviousLocationAreaURL *string
	caughtPokemon           map[string]pokeapi.Pokemon
}

func main() {
	httpTimeout := 5 * time.Second
	cacheTTL := 5 * time.Minute
	cfg := Config{
		pokeapiClient: pokeapi.NewClient(httpTimeout, cacheTTL),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
