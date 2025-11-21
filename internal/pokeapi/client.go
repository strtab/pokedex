package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	mapNext string
	mapPrev string
)

type locationAreaResp struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(isNext bool) error {
	url := "https://pokeapi.co/api/v2/location-area/?limit=20"

	if isNext && mapNext != "" {
		url = mapNext
	} else if !isNext && mapPrev != "" {
		url = mapPrev
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

	var resp locationAreaResp
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
