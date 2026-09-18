package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(configPath)
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil

}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(homePath, configFileName)
	return path, nil
}

func SetUser(config *Config, username string) error {
	config.CurrentUserName = username
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
