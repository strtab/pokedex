package main

import (
	"bufio"
	"fmt"
	"github.com/strtab/pokedex/internal/commands"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")

		if !sc.Scan() {
			commands.CmdExit()
		}
		out := cleanInput(sc.Text())

		if len(out) == 0 {
			continue
		}

		if cmd, ok := commands.Register[out[0]]; ok {
			if err := cmd.Callback(); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
