package main

import "fmt"

func commandMapB([]string) error {
	prevUrl := config.previous

	if prevUrl == "" {
		err := fmt.Errorf("There are no previous locations")
		println(err.Error())
		return err
	}

	locations, err := getLocations(prevUrl)

	for _, value := range locations {
		println(value)
	}

	return err
}
