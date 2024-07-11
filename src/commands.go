package main

import (
	"fmt"
	"os"
)

type cliCommands struct {
	name        string
	description string
	callback    func() error
}

func commands() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Lists available commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Here are the available commands: ")

	availableCommands := commands()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}

func commandExit() error {
	fmt.Println("Closing pokedex")
	os.Exit(0)
	return nil
}
