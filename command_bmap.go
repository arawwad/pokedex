package main

import "fmt"

func commandMapB() error {
	prevUrl := config.previous

	if prevUrl == "" {
		err := fmt.Errorf("There are no previous locations")
		println(err.Error())
		return err
	}

	return getLocations(prevUrl, &config)
}