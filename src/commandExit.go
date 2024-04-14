package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println()
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}
