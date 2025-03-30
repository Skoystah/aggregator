package config

import (
	"encoding/json"
	"os"
)

const configFile = ".gatorconfig.json"

type Config struct {
	URL             string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {

	cfgFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.ReadFile(cfgFilePath)

	cfg := Config{}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c *Config) SetUser(userName string) error {

	cfgFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	c.CurrentUserName = userName

	//Perhaps create helper function and/or use NewEncoder instead of Marshal + Write file
	cfgFile, err := json.Marshal(c)
	err = os.WriteFile(cfgFilePath, cfgFile, 0777)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + configFile, nil
}
