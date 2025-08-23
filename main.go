package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
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

type LocResult struct {
	Name string
	URL  string
}
type LocArea struct {
	Count      int
	Next       *string     // can be null
	Previous   *string     // can be null
	LocResults []LocResult `json:"results"`
}

/** displays the names of 20 location areas in the Pokemon world */
func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var locArea LocArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locArea)
	if err != nil {
		return err
	}

	for _, r := range locArea.LocResults {
		fmt.Println(r.Name)
	}
	return nil
}

/** It's similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back. */
func commandMapBack() error {
	// todo: implement this using config struct of next and prev to hold the next and previous url

	// If you're on the first "page" of results, this command should just print "you're on the first page"
	return nil
}

/** contain the Next and Previous URLs that you'll need to paginate */
type config struct {
	next *string
	prev *string
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *config
}

func main() {
	commandMap := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
			config:      &config{},
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			config:      &config{},
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
			config:      &config{},
		},
		"mapb": {
			name:        "mapb",
			description: "It's similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.",
			callback:    commandMapBack,
			config:      &config{},
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
