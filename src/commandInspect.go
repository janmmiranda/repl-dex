package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, pokemonName string) error {
	if pokemonName == "" {
		return errors.New("pass in a pokemon's name to inspect")
	}
	pokemon, err := cfg.pokemonBox.pokemonGet(pokemonName)
	if err != nil {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %v\n", t.Type.Name)
	}

	return nil
}
