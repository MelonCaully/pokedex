package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)

	if err != nil {
		return err
	}

	chance := rand.IntN(628)
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if chance > pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped\n", pokemon.Name)
	}

	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
