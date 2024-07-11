package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		scanner.Scan()
		text := scanner.Text()

		cleanedInput := cleanUserInput(text)
		if len(cleanedInput) == 0 {
			continue
		}

		commandInput := cleanedInput[0]
		availableCommands := commands()

		command, ok := availableCommands[commandInput]

		if !ok {
			fmt.Println("That command is not available. Try 'help' for further assistance.")
			continue
		}

		err := command.callback()

		if err != nil {
			fmt.Println(err)
		}

	}
}

func cleanUserInput(s string) []string {
	lower := strings.ToLower(s)
	words := strings.Fields(lower)

	return words
}
