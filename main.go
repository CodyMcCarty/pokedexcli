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

func main() {
	// Create support for a simple REPL
	// Wait for user input using bufio.NewScanner (this blocks the code and waits for input, once the user types something and presses enter, the code continues and the input is available in the returned bufio.Scanner)
	scanner := bufio.NewScanner(os.Stdin)
	// Start an infinite for loop. This loop will execute once for every command the user types in (we don't want to exit the program after just one command)
	for {
		// Use fmt.Print to print the prompt Pokedex > without a newline character
		fmt.Print("Pokedex > ")
		// Use the scanner's .Scan and .Text methods to get the user's input as a string
		scanner.Scan()
		str := scanner.Text()
		// Clean the user's input string
		strs := cleanInput(str)
		// Capture the first "word" of the input and use it to print: Your command was: <first word>
		fmt.Println("Your command was:", strs[0])
	}
	// Test your program. Here's my example session:
	// Run the CLI again and tee the output (copies the stdout) to a new file called repl.log (and .gitignore the log).
}
