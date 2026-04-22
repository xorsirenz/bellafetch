package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Ascii   string          `json:"Ascii"`
	Modules map[string]bool `json:"Modules"`
}

func LoadConfig() Config {
	configFile, err := configDirExists()
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	file, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("cannot read config: %v", err)
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Cannot unmarshal json: %v", err)
	}

	return config
}

func configDirExists() (string, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return "", fmt.Errorf("Error failed to find user config direcoty: %v", err)
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := createDefaultConfig(configPath); err != nil {
			return "", err
		}
	}
	return configPath, err
}

func getConfigPath() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("Failed to find user config directory: %w", err)
	}
	return filepath.Join(userConfigDir, "bellafetch", "config"), err
}

func createDefaultConfig(configPath string) error {
	configDir := filepath.Dir(configPath)

	configJson := Config{
		Ascii: "none",
		Modules: map[string]bool{
			"Host":       true,
			"PrettyName": true,
			"Kernel":     true,
			"Uptime":     true,
			"Packages":   true,
			"Shell":      true,
			"Terminal":   true,
			"WM":         true,
			"Cpu":        true,
			"Gpu":        true,
			"DiskSpace":  true,
			"Memory":     true,
		},
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("Failed to create config directory: %w", err)
	}

	defaultConfigData, err := json.MarshalIndent(configJson, "", "  ")
	if err != nil {
		return fmt.Errorf("Failed to Marshal config file: %w", err)
	}
	if err := os.WriteFile(configPath, defaultConfigData, 0644); err != nil {
		return fmt.Errorf("Failed to write default config file: %w", err)
	}
	return err
}
