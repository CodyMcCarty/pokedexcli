# Pokedex cli

A simple **Pokédex CLI tool** written in Go.  
This project is part of a boot.dev assignment to practice Go programming, HTTP requests, caching, and building a REPL (Read–Eval–Print Loop).

## Features

- Interactive REPL prompt (`Pokedex >`) for commands
- `help` — list available commands
- `exit` — quit the program
- `...` — others that hit the poke API or do something with data
- In-memory caching layer for API responses with a concurrent automatic expiration (to reduce network calls)

## Getting Started

### Prerequisites
- [Go 1.25+](https://go.dev/dl/)

### Run
Clone the repo and run directly:

```bash
git clone https://github.com/CodyMcCarty/pokedexcli.git
cd pokedexcli
go run .
```
### Example Session
```bash
Pokedex > help
Available commands:
  help    Show available commands
  map     Display a list of Pokémon location areas
  mapb    Go back to the previous page
  exit    Exit the program
  ...

Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > mapb
you\'re on the first page

Pokedex > exit
Goodbye!
```

### Testing
Run the tests with `$ go test ./...`

## Caching
API responses are cached in memory for a configurable interval. Old entries are automatically reaped by a background goroutine. This improves responsiveness when paging back and forth.

## Notes
This is an educational assignment, not a production-ready tool. The focus is on learning Go basics: REPL loops, concurrency, HTTP clients, and simple caching.  Contributions/PRs aren’t expected — this repo is mainly for practice. An API key is not necessary for the PokeAPI

## Future Extensions
- Update the CLI to support the "up" arrow to cycle through previous commands
- Simulate battles between pokemon
- Add more unit tests
- Refactor your code to organize it better and make it more testable
- Keep pokemon in a "party" and allow them to level up
- Allow for pokemon that are caught to evolve after a set amount of time
- Persist a user's Pokedex to disk so they can save progress between sessions
- Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- Random encounters with wild pokemon
- Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon


