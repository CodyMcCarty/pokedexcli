package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/** displays the names of 20 location areas in the Pokemon world */
func commandMap() error {
	fullUrl := pokeURL + "location-area/"
	if config.next != nil {
		fullUrl = *config.next
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var page LocArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&page)
	if err != nil {
		return err
	}

	//locArea.Previous // this is the previous url, what do i do with it?

	for _, r := range page.LocResults {
		fmt.Println(r.Name)
	}

	config.next = page.Next
	config.prev = page.Previous
	return nil
}

/** It's similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back. */
func commandMapBack() error {
	if config.prev == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(*config.prev)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var page LocArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&page)
	if err != nil {
		return err
	}

	for _, r := range page.LocResults {
		fmt.Println(r.Name)
	}

	config.next = page.Next
	config.prev = page.Previous

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
