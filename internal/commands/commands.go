package commands

import (
	"fmt"
	"github.com/strtab/pokedex/internal/pokeapi"
	"os"
)

func CmdMap(args ...string) error {
	return pokeapi.GetLocationAreas(true)
}

func CmdMapb(args ...string) error {
	return pokeapi.GetLocationAreas(false)
}

func CmdPokedex(args ...string) error {
	return pokeapi.GetCaughtPokemons()
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
