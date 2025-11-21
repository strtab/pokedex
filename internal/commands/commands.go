package commands

import (
	"fmt"
	"github.com/strtab/pokedex/internal/pokeapi"
	"os"
)

var Register map[string]cliCmd

func init() {
	Register = map[string]cliCmd{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CmdExit,
		},
		"map": {
			Name:        "map",
			Description: "Display the next 20 locations",
			Callback:    CmdMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display the previous 20 locations",
			Callback:    CmdMapb,
		},
		"help": {
			Name:        "help",
			Description: "Displays the help message",
			Callback:    CmdHelp,
		},
	}
}

type cliCmd struct {
	Name        string
	Description string
	Callback    func() error
}

func CmdMap() error {
	return pokeapi.GetLocationAreas(true)
}

func CmdMapb() error {
	return pokeapi.GetLocationAreas(false)
}

func CmdExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CmdHelp() error {
	fmt.Println("Usage:")
	for _, cmd := range Register {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}
