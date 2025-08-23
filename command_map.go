package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/** displays the names of 20 location areas in the Pokemon world */
func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var locArea LocArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locArea)
	if err != nil {
		return err
	}

	//locArea.Previous // this is the previous url, what do i do with it?

	for _, r := range locArea.LocResults {
		fmt.Println(r.Name)
	}
	return nil
}

/** It's similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back. */
func commandMapBack() error {
	// todo: implement this using config struct of next and prev to hold the next and previous url

	// If you're on the first "page" of results, this command should just print "you're on the first page"
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
