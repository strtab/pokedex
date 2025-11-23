package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
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

	if cacheData, exsist := cache.Get(url); exsist {
		return readAreas(cacheData)
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

	if cacheData, exsist := cache.Get(url); exsist {
		return readPokemonEncounters(cacheData)
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

func CatchPokemon(name string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)

	if cacheData, exsist := cache.Get(url); exsist {
		return tryCatchPokemon(cacheData)
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
	return tryCatchPokemon(body)
}

func tryCatchPokemon(body []byte) error {
	var resp pokemon
	if err := json.Unmarshal(body, &resp); err != nil {
		return fmt.Errorf("Pokemon not found")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", resp.Name)

	time.Sleep(time.Second)
	if rand.Intn(resp.BaseExperience/2) < resp.BaseExperience/6 {
		fmt.Printf("%s was caught!\n", resp.Name)
	} else {
		fmt.Printf("%s escaped!\n", resp.Name)
	}

	return nil
}

func InspectPokemon(name string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)

	if cacheData, exsist := cache.Get(url); exsist {
		return pokemonDataRead(cacheData)
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
	return pokemonDataRead(body)
}

func pokemonDataRead(body []byte) error {
	var resp pokemon
	if err := json.Unmarshal(body, &resp); err != nil {
		return fmt.Errorf("Pokemon not found")
	}

	fmt.Println("Name:", resp.Name)
	fmt.Println("Height:", resp.Height)
	fmt.Println("Weight:", resp.Weight)
	fmt.Println("Stats:")
	for _, value := range resp.Stats {
		fmt.Printf(" - %v: %d\n", value.Stat.Name, value.BaseStat)
	}
	fmt.Println("Types:")
	for _, value := range resp.Types {
		fmt.Printf(" - %v\n", value.Type.Name)
	}
	return nil
}
