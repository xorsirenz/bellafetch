package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Modules []Modules `json:"Modules"`
}

type Modules struct {
	Host       bool `json:"Host"`
	PrettyName bool `json:"PrettyName"`
	Kernel     bool `json:"Kernel"`
	Uptime     bool `json:"Uptime"`
	Package    bool `json:"Packages"`
	WM         bool `json:"WM"`
	Cpu        bool `json:"Cpu"`
	Gpu        bool `json:"Gpu"`
	DiskSpace  bool `json:"DiskSpace"`
	Memory     bool `json:"Memory"`
}

func LoadConfig() (map[string]bool, error) {
	configFile, err := configDirExists()
	if err != nil {
		fmt.Errorf("cannot load config: %v", err)
		os.Exit(1)
	}

	file, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var configData struct {
		Modules []map[string]bool `json:"Modules"`
	}

	err = json.Unmarshal(file, &configData)
	if err != nil {
		fmt.Errorf("Cannot unmarshal json: %w", err)
		os.Exit(1)
	}

	return configData.Modules[0], nil
}

func configDirExists() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("Failed to find user config directory: %w", err)
	}
	bellafetchConfigPath := filepath.Join(configDir, "bellafetch")

	if err := os.MkdirAll(bellafetchConfigPath, 0755); err != nil {
		return "", fmt.Errorf("Failed to create config directory: %w", err)
	}

	configJson := Config{
		Modules: []Modules{
			{
				Host:       true,
				PrettyName: true,
				Kernel:     true,
				Uptime:     true,
				Package:    true,
				WM:         false,
				Cpu:        true,
				Gpu:        true,
				DiskSpace:  true,
				Memory:     true,
			},
		},
	}

	configFilePath := filepath.Join(bellafetchConfigPath, "config")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		defaultConfigData, err := json.MarshalIndent(configJson, "", "  ")
		if err != nil {
			return "", fmt.Errorf("Failed to Marshal config file: %w", err)
		}
		if err := os.WriteFile(configFilePath, defaultConfigData, 0644); err != nil {
			return "", fmt.Errorf("Failed to write default config file: %w", err)
		}
	}
	return configFilePath, err
}
