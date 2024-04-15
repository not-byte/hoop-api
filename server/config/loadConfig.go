package config

import (
	"encoding/json"
	"fmt"
	"os"
	"tournament_api/server/types"
)

func LoadConfig() (*types.AppConfig, error) {
	var config types.AppConfig

	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	//this might be later changed to relative path
	configFile, err := os.Open(fmt.Sprintf("server/config/config.%s.json", env))

	if err != nil {
		return nil, err
	}

	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
