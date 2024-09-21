package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func([]string) error
}

const prompt = "Pokedex > "

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := strings.Fields(strings.TrimSpace(scanner.Text()))

		if len(input) == 0 {
			continue
		}

		command, ok := commands[input[0]]

		if ok {
			command.callback(input[1:])
		} else {
			fmt.Println("Command not found")
		}
	}
}
