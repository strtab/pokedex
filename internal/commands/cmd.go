package cmd

import (
	"fmt"
)

var register map[string]cliCmd

func init() {
	register = map[string]cliCmd{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    cmdExit,
		},
		"map": {
			name:        "map",
			description: "Display the map",
			callback:    cmdMap,
		},
		"help": {
			name:        "help",
			description: "Displays the help message",
			callback:    cmdHelp,
		},
	}
}

type cliCmd struct {
	name        string
	description string
	callback    func() error
}

func cmdExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cmdHelp() error {
	fmt.Println("Usage:")
	for _, cmd := range register {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
