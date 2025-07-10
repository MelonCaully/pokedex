package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MelonCaully/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, ok := getCommands()[commandName]

		if ok {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown commands")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the Location Areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the Location Areas of the previous url",
			callback:    commandMapb,
		},
	}
}
