package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/** The purpose of this function will be to split the user's input into "words" based on whitespace. It should also lowercase the input and trim any leading or trailing whitespace */
func cleanInput(text string) []string {
	var result []string

	text = strings.ToLower(text)
	result = strings.Fields(text)

	return result
}

/** callback for the exit command */
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	// todo: You can dynamically generate the "usage" section by iterating over my registry of commands. That way the help command will always be up-to-date with the available commands.
	fmt.Print(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	commandMap := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		str := scanner.Text()
		//strs := cleanInput(str)

		command, ok := commandMap[str]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback()
		if err != nil {
			continue
		}
	}
}
