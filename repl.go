package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

const prompt = "Pokedex > "

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		scanner.Scan()
		commandName := scanner.Text()
		command, ok := commands[commandName]

		if ok {
			command.callback()
		} else {
			fmt.Println("Command not found")
		}
	}
}
