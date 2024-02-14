package main

import (
	"errors"
	"fmt"
)

func map_cmd(cfg *config, args ...string) error {

	locArea, err := cfg.pokeapiClient.GetLocations(cfg.next)
	if err != nil {
		return err
	}
	for _, result := range locArea.Results {
		fmt.Println(result.Name)
	}
	cfg.previous = locArea.Previous
	cfg.next = locArea.Next
	return nil
}

func mapb_cmd(cfg *config, args ...string) error {
	if cfg.previous == nil {
		return errors.New("already on first page, no previous available")
	}
	locArea, err := cfg.pokeapiClient.GetLocations(cfg.previous)
	if err != nil {
		return err
	}
	for _, result := range locArea.Results {
		fmt.Println(result.Name)
	}
	cfg.next = locArea.Next
	cfg.previous = locArea.Previous
	return nil
}
