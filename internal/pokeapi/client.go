package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/strtab/pokedex/internal/pokecache"
)

var (
	mapNext string
	mapPrev string
	cache   pokecache.Cache
)

func init() {
	cache = *pokecache.NewCache(time.Minute + 30)
}

func GetLocationAreas(isNext bool) error {
	url := "https://pokeapi.co/api/v2/location-area/?limit=20"

	if isNext && mapNext != "" {
		url = mapNext
	} else if !isNext && mapPrev != "" {
		url = mapPrev
	}

	if val, exsist := cache.Get(url); exsist {
		return readAreas(val)
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	cache.Add(url, body)

	return readAreas(body)
}

func readAreas(body []byte) error {
	var resp locationAreas
	if err := json.Unmarshal(body, &resp); err != nil {
		return err
	}

	mapNext = resp.Next
	mapPrev = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func ExploreLocation(location string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)

	fmt.Printf("Exploring %s...\n", location)

	if locationArea, exsist := cache.Get(url); exsist {
		return readPokemonEncounters(locationArea)
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	cache.Add(url, body)
	return readPokemonEncounters(body)
}

func readPokemonEncounters(body []byte) error {
	var resp locationArea
	if err := json.Unmarshal(body, &resp); err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, v := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", v.Pokemon.Name)
	}

	return nil
}
