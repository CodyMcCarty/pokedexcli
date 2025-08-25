package main

import (
	"time"

	"github.com/CodyMcCarty/pokedexcli/internal/pokeapi"
)

const pokeURL = "https://pokeapi.co/api/v2/"

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
