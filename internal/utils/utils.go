package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadConfig() (Configuration, error) {
	file, err := os.Open("data/config.json")
	if err != nil {
		return Configuration{}, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var config Configuration
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return Configuration{}, fmt.Errorf("error decoding JSON: %v", err)
	}
	return config, nil
}
