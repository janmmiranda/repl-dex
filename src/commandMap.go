package main

import (
	"fmt"
)

func commandMapf(cfg *config) error {
	fmt.Println()
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsUrl)
	if err != nil {
		return err
	}
	cfg.nextLocationsUrl = locationsResp.Next
	cfg.previousLocationsUrl = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	fmt.Println()
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = locationsResp.Next
	cfg.previousLocationsUrl = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
