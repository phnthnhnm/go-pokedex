package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/phnthnhnm/go-pokedex/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	Next     string
	Previous string
	Pokedex  map[string]api.Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore a specific location area and list its Pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon by name",
			callback:    commandCatch,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		Next:    "https://pokeapi.co/api/v2/location-area/",
		Pokedex: make(map[string]api.Pokemon),
	}
	rand.Seed(time.Now().UnixNano()) // Seed random number generator
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) > 0 {
			commandName := words[0]
			args := words[1:]
			if command, found := getCommands()[commandName]; found {
				if err := command.callback(cfg, args); err != nil {
					fmt.Printf("Error: %s\n", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	words := strings.Fields(text)

	return words
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config, args []string) error {
	if cfg.Next == "" {
		fmt.Println("No more locations to display.")
		return nil
	}

	data, err := api.FetchLocationAreas(cfg.Next)
	if err != nil {
		return err
	}

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous
	return nil
}

func commandMapBack(cfg *config, args []string) error {
	if cfg.Previous == "" {
		fmt.Println("You're on the first page.")
		return nil
	}

	data, err := api.FetchLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous
	return nil
}

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		fmt.Println("Usage: explore <location-area-name>")
		return nil
	}

	locationName := args[0]
	fmt.Printf("Exploring %s...\n", locationName)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", locationName)
	data, err := api.FetchLocationAreaDetails(url)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range data.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		fmt.Println("Usage: catch <pokemon-name>")
		return nil
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName)
	pokemon, err := api.FetchPokemonDetails(url)
	if err != nil {
		return fmt.Errorf("failed to fetch Pokemon details: %w", err)
	}

	catchChance := 50 + (100-pokemon.BaseExperience)/2
	if catchChance > 95 {
		catchChance = 95 // Cap the maximum chance at 95%
	} else if catchChance < 5 {
		catchChance = 5 // Ensure a minimum chance of 5%
	}

	if rand.Intn(100) >= catchChance {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	// Add Pokemon to user's Pokedex
	cfg.Pokedex[pokemonName] = *pokemon
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}
