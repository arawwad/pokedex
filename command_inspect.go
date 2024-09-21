package main

import "fmt"

func commandInspect(params []string) error {

	if len(params) == 0 {
		fmt.Println("Can't inspect nothing!")
		return fmt.Errorf("user input was empty")
	}

	name := params[0]
	value, ok := pokedex[name]
	if !ok {
		println("you have not caught that pokemon")
	} else {
		println("Name: ", value.name)
		println("Height: ", value.height)
		println("Weight: ", value.weight)
		println("stats:")
		for statName, stat := range value.stats {
			println("   -", statName, ": ", stat)
		}
		println("Types:")
		for _, pType := range value.types {
			println("   -", pType)
		}
	}

	return nil
}
