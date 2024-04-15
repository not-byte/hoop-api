package config

import (
	"encoding/json"
	"fmt"
	"go/types"
	"log"
	"os"
)

func LoadConfig() (*types.Config, error) {
	var config types.Config

	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	cwd, _ := os.Getwd()
	log.Println("Current Working Directory:", cwd)

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
