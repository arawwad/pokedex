package main

import "fmt"

func commandExplore(params []string) error {
	if len(params) == 0 {
		fmt.Println("Can't explore nothing!")
		return fmt.Errorf("user input was empty")
	}

	pokemons, err := getPokemonInLocation(params[0])

	println("Found Pokemon:")
	for _, value := range pokemons {
		pokeLine := fmt.Sprintf("    - %s", value)
		println(pokeLine)
	}

	return err
}
