package main

func commandPokedex([]string) error {

	println("Your Pokedex:")
	for key, _ := range pokedex {
		println(" -", key)
	}

	return nil
}
