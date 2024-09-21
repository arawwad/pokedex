package main

import "fmt"

func commandExplore(params []string) error {
	if len(params) == 0 {
		fmt.Println("Can't explore nothing!")
		return fmt.Errorf("user input was empty")
	}

	err := getLocation(params[0])

	return err
}
