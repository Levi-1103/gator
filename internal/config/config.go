package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg Config) SetUser(username string) error {
	cfg.CurrentUserName = username

	err := write(cfg)
	if err != nil {
		return err
	}

	return nil

}

func write(cfg Config) error {
	configLocation, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonConfig, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configLocation, jsonConfig, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {

	configLocation, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configLocation)
	if err != nil {
		return Config{}, err
	}

	var config Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error decoding json: %w", err)
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := homeDir + "/" + configFileName

	return configPath, nil
}
