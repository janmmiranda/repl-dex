package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func commandExit(cfg *config, area string) error {
	if f, err := os.Create(".repl_history"); err == nil {
		cfg.line.WriteHistory(f)
		cfg.line.Close()
		f.Close()
	}

	if f, err := os.Create(pokemonFilename); err == nil {
		defer f.Close()

		encoder := json.NewEncoder(f)
		err = encoder.Encode(cfg.pokemonBox.box)
		if err != nil {
			return err
		}

	}

	fmt.Println()
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}
