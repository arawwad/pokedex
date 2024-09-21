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

type serviceConfig struct {
	next     string
	previous string
}

func getLocations(url string, config *serviceConfig) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	response := &locationsResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		return err
	}

	config.next = response.Next
	config.previous = response.Previous

	for _, value := range response.Results {
		println(value.Name)
	}

	return nil
}
