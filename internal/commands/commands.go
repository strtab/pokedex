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
		"explore": {
			Name:        "explore",
			Description: "Explore location",
			Callback:    CmdExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon",
			Callback:    CmdCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a pokemon",
			Callback:    CmdInspect,
		},
		"help": {
			Name:        "help",
			Description: "Displays the help message",
			Callback:    CmdHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CmdExit,
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

func CmdCatch(args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("Usage: catch <pokemon>")
	}

	return pokeapi.CatchPokemon(args[1])
}

func CmdInspect(args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("Usage: inspect <pokemon>")
	}

	return pokeapi.InspectPokemon(args[1])
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
