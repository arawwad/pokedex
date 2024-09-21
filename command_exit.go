package main

import (
	"fmt"
	"os"
)

func commandExit([]string) error {
	fmt.Println("Exiting Pokedex...")
	os.Exit(0)
	return nil
}
