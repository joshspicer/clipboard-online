package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

const ConfigFile = "config.json"
const LogFile = "log.txt"

// Config represents configuration for applicaton
type Config struct {
	Port           string       `json:"port"`
	LogLevel       logrus.Level `json:"logLevel"`
	TempDir        string       `json:"tempDir"`
	ReserveHistory bool         `json:"reserveHistory"`
}

// DefaultConfig is a default configuration for application
var DefaultConfig = Config{
	Port:           "8086",
	LogLevel:       log.WarnLevel,
	TempDir:        "./temp",
	ReserveHistory: false,
}

// DefaultConfigCopy returns a deep copy of DefaultConfig
func DefaultConfigCopy() *Config {
	config := Config{
		Port:           DefaultConfig.Port,
		LogLevel:       DefaultConfig.LogLevel,
		TempDir:        DefaultConfig.TempDir,
		ReserveHistory: DefaultConfig.ReserveHistory,
	}
	return &config
}

func loadConfig(path string) (*Config, error) {
	if isExistFile(path) {
		return loadConfigFromFile(path)
	}
	if err := createConfigFile(path); err != nil {
		return nil, err
	}
	return &DefaultConfig, nil
}

func loadConfigFromFile(path string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := DefaultConfigCopy()
	if err := json.Unmarshal(configBytes, config); err != nil {
		return nil, err
	}
	return config, nil
}

func createConfigFile(path string) error {
	defaultConfigJSON, err := json.MarshalIndent(DefaultConfig, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, []byte(defaultConfigJSON), 0744); err != nil {
		return err
	}
	return nil
}

func isExistFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
