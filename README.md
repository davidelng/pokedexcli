# Pokedex CLI

A Pokemon REPL written in Go. Build and launch with `go build -o out && ./out`. It calls the [PokeApi](https://pokeapi.co/) but caches the responses for performance

## Commands

- `help` displays a help message
- `map` show next map locations
- `mapb` show previous map locations
- `explore {location_area}` list the pokemon in a location area
- `catch {pokemon_name}` attempt to catch a pokemon and add it to your pokedex
- `inspect {pokemon_name}` show stat of a caught pokemon
- `pokedex` show your pokedex
- `exit` exit the Pokedex

## Ideas to extend the project

- Update the CLI to support the "up" arrow to cycle through previous commands
- Simulate battles between pokemon
- Keep pokemon in a "party" and allow them to level up
- Allow for pokemon that are caught to evolve after a set amount of time
- Persist a user's Pokedex to disk so they can save progress between sessions
- Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- Random encounters with wild pokemon
- Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon
