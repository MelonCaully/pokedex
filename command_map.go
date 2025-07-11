package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)

	if err != nil {
		return err
	}

	cfg.nextLocationsURL = &locationsResp.Next
	if locationsResp.Previous != nil {
		if prevURLStr, ok := locationsResp.Previous.(string); ok {
			cfg.prevLocationsURL = &prevURLStr
		} else {
			cfg.prevLocationsURL = nil
		}
	} else {
		cfg.prevLocationsURL = nil
	}

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)

	if err != nil {
		return err
	}

	cfg.nextLocationsURL = &locationsResp.Next
	if locationsResp.Previous != nil {
		if prevURLStr, ok := locationsResp.Previous.(string); ok {
			cfg.prevLocationsURL = &prevURLStr
		} else {
			cfg.prevLocationsURL = nil
		}
	} else {
		cfg.prevLocationsURL = nil
	}

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
