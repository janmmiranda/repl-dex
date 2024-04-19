package main

import (
	"time"

	"github.com/janmmiranda/repl-dex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokemonBox:    pokeBoxCreate(),
	}
	startRepl(cfg)
}
