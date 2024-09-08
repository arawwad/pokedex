package main

// Initialize all commands using an init function
var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"help": {name: "help", description: "Display a help message", callback: commandHelp},
		"exit": {name: "exit", description: "Exit the Pokedex", callback: commandExit},
	}
}
