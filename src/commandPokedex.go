package main

import (
	"fmt"
)

func commandPokedex(cfg *config, pokemonName string) error {
	fmt.Println("Your Pokedex:")
	for _, v := range cfg.pokemonBox.box {
		fmt.Printf("  - %s\n", v.Pokemon.Name)
	}
	return nil
}
