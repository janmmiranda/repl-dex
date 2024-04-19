package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, pokemonName string) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	pokemon, err := cfg.pokeapiClient.PokemonFetch(pokemonName)
	if err != nil {
		return err
	}
	r := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if r > 40 {
		fmt.Printf("%s escaped! \n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	cfg.pokemonBox.pokemonAdd(pokemon)
	return nil
}
