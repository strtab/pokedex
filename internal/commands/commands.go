package commands

import (
	"fmt"
	"github.com/strtab/pokedex/internal/pokeapi"
	"os"
)

var Register map[string]cliCmd

type cliCmd struct {
	Name        string
	Description string
	Callback    func(args ...string) error
}

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
		"explore": {
			Name:        "explore",
			Description: "Explore location",
			Callback:    CmdExplore,
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

func CmdMap(args ...string) error {
	return pokeapi.GetLocationAreas(true)
}

func CmdMapb(args ...string) error {
	return pokeapi.GetLocationAreas(false)
}

func CmdExplore(args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("Usage: explore <location>")
	}
	return pokeapi.ExploreLocation(args[1])
}

func CmdExit(args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CmdHelp(args ...string) error {
	fmt.Println("Usage:")
	for _, cmd := range Register {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}
