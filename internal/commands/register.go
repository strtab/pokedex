package commands

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
		"pokedex": {
			Name:        "pokedex",
			Description: "Display caught pokemons",
			Callback:    CmdPokedex,
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
