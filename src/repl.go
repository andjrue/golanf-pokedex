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
		fmt.Println("Echoing: ", cleanedInput)
		// command := cleanedInput[0]
		// availableCommands := commands() ** Will get to this later

	}
}

func cleanUserInput(s string) []string {
	lower := strings.ToLower(s)
	words := strings.Fields(lower)

	return words
}
