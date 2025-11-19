package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var mapNext string

func cmdMap() error {
	url := "https://pokeapi.co/api/v2//location-area/"
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var cityArea string
	json.Unmarshal(body, &cityArea)
	return nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")

		if !sc.Scan() {
			cmdExit()
		}
		out := cleanInput(sc.Text())

		if len(out) == 0 {
			continue
		}

		if cmd, ok := register[out[0]]; ok {
			if err := cmd.callback(); err != nil {
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
