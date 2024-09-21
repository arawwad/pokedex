package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(params []string) error {

	if len(params) == 0 {
		fmt.Println("Can't catch nothing!")
		return fmt.Errorf("user input was empty")
	}

	name := params[0]
	if _, ok := pokedex[name]; ok {
		println("You already caught ", name, "!")
		return nil
	}

	println("Throwing a Pokeball at ", name, " ...")

	pokemon, err := getPokemon(name)
	if err != nil {
		return err
	}

	chance := float32(pokemon.baseExperience) / 400.0

	if rand.Float32() > chance {
		println(name, " was caught!")
		pokedex[name] = pokemon
	} else {
		println(name, " escaped!")
	}

	return nil
}
