package main

func commandMap([]string) error {
	locations, err := getLocations(config.next)

	for _, value := range locations {
		println(value)
	}
	return err
}
