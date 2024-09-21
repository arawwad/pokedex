package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationsResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type locationResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func getFromCacheOrServer(url string) ([]byte, error) {
	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			return nil, err
		}

		cache.Add(url, body)
	}

	return body, nil
}

func getLocations(url string) ([]string, error) {
	body, err := getFromCacheOrServer(url)
	if err != nil {
		return []string{}, err
	}

	response := &locationsResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		return []string{}, err
	}

	config.next = response.Next
	config.previous = response.Previous

	var locations []string
	for _, value := range response.Results {
		locations = append(locations, value.Name)
	}

	return locations, nil
}

func getPokemonInLocation(name string) ([]string, error) {
	body, err := getFromCacheOrServer(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name))

	if err != nil {
		return []string{}, err
	}

	response := &locationResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		return []string{}, err
	}

	var pokemons []string
	for _, value := range response.PokemonEncounters {
		pokemons = append(pokemons, value.Pokemon.Name)
	}

	return pokemons, nil
}
