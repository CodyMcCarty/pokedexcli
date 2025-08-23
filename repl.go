package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		commandWords := cleanInput(userInput)
		commandName := commandWords[0]

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

/** contain the Next and Previous URLs that you'll need to paginate */
type configuration struct {
	next *string
	prev *string
}

var config = &configuration{}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
			//config:      &config{},
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			//config:      &config{},
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
			//config:      &config{},
		},
		"mapb": {
			name:        "mapb",
			description: "It's similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.",
			callback:    commandMapBack,
			//config:      &config{},
		},
	}
}

/** The purpose of this function will be to split the user's input into "words" based on whitespace. It should also lowercase the input and trim any leading or trailing whitespace */
func cleanInput(text string) []string {
	var result []string

	text = strings.ToLower(text)
	result = strings.Fields(text)

	return result
}
