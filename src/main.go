package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/janmmiranda/repl-dex/internal/pokeapi"
	"github.com/peterh/liner"
)

const pokemonFilename = "pokemon.json"

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	l := liner.NewLiner()
	defer func() {
		l.Close()
		if err := recover(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}()
	cfg := &config{
		pokeapiClient: pokeClient,
		pokemonBox:    pokeBoxCreate(),
		line:          l,
	}
	existingPokemon, err := readPokemonFromFile(pokemonFilename)
	if err == nil {
		cm := pokeboxCaughtMapSet(existingPokemon)
		cfg.pokemonBox.box = existingPokemon
		cfg.pokemonBox.caughtMap = cm
	}
	startRepl(cfg)
}

func pokeboxCaughtMapSet(pokemonMap map[uuid.UUID]pokeapi.Wpokemon) map[string][]uuid.UUID {
	cm := make(map[string][]uuid.UUID)
	for k, v := range pokemonMap {
		if val, ok := cm[v.Pokemon.Name]; ok {
			val = append(val, k)
		} else {
			tempSlice := make([]uuid.UUID, 0)
			tempSlice = append(tempSlice, k)
			val = tempSlice
		}
	}
	return cm
}

func readPokemonFromFile(filename string) (map[uuid.UUID]pokeapi.Wpokemon, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return make(map[uuid.UUID]pokeapi.Wpokemon), err
	}

	var data map[uuid.UUID]pokeapi.Wpokemon
	err = json.Unmarshal(file, &data)
	if err != nil {
		return make(map[uuid.UUID]pokeapi.Wpokemon), err
	}
	return data, nil
}
