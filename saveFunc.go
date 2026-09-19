package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func saveFunction(conf *config) error {

	data, err := json.Marshal(conf.Pokedex)
	if err != nil {
		return fmt.Errorf("error saving data")
	}

	err = os.WriteFile("saves/save_data.json", data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func loadFunction(conf *config) error {
	data, err := os.ReadFile("saves/save_data.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &conf.Pokedex)
	if err != nil {
		return err
	}
	return err

}
