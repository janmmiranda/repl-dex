package main

import (
	"fmt"
)

func commandExplore(cfg *config, area string) error {
	encountersRes, err := cfg.pokeapiClient.ExploreLocation(area)
	if err != nil {
		return err
	}

	for _, pokemon := range encountersRes.PokemonEncounters {
		fmt.Println(" - ", pokemon.Pokemon.Name)
	}

	return nil
}
