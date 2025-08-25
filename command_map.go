package main

import (
	"errors"
	"fmt"
)

/** displays the names of 20 location areas in the Pokemon world */
func commandMapFwd(cfg *config) error {

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

/** It's similar to the map command, however, instead of displaying the nextLocationsURL 20 locations, it displays the previous 20 locations. It's a way to go back. */
func commandMapBack(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

type LocArea struct {
	Count      int
	Next       *string     // can be null
	Previous   *string     // can be null
	LocResults []LocResult `json:"results"`
}

type LocResult struct {
	Name string
	URL  string
}
