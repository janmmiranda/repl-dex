package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, area string) error {
	fmt.Println()
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}
