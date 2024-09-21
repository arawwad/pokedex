package main

// Initialize all commands using an init function
var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"help":    {name: "help", description: "Display a help message", callback: commandHelp},
		"exit":    {name: "exit", description: "Exit the Pokedex", callback: commandExit},
		"map":     {name: "map", description: "Get the next locations", callback: commandMap},
		"mapb":    {name: "mapb", description: "Get the previous locations", callback: commandMapB},
		"explore": {name: "explore", description: "Explore Location", callback: commandExplore},
		"catch":   {name: "catch", description: "Catch a pokemon by name", callback: commandCatch},
	}
}
