package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(conf *Config) error {
	url := conf.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}
	locationArea := getLocationArea(url, conf)
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapb(conf *Config) error {
	url := conf.Previous
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locationArea := getLocationArea(url, conf)
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func getLocationArea(url string, conf *Config) LocationArea {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)

	}
	if err != nil {
		fmt.Println(err)
	}
	var locationArea LocationArea
	if err := json.Unmarshal(body, &locationArea); err != nil {
		fmt.Printf("Error unmarshalling JSON: %v", err)
	}
	conf.Next = locationArea.Next
	conf.Previous = locationArea.Previous
	return locationArea
}

type LocationArea struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
